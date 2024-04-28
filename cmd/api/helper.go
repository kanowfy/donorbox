package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data map[string]interface{}, headers http.Header) error {
	encoded, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	encoded = append(encoded, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(encoded)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dest interface{}) error {
	// limit request body to maximum of 1MB
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dest)
	if err != nil {
		// if any error occurred during decoding, start the triage
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return fmt.Errorf("body contains badly formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		// if the request body contains an unknown key, return an error instead of silently discard it
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		// json.InvalidUnmarshalError will be return if we pass a non-nil pointer to Decode()
		// we catch this and panic
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	// call decode again to an empty anonymous struct to check if the request body contains additional junk data
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func readString(qs url.Values, key, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	return s
}

func readInt(qs url.Values, key string, defaultValue int) (int, error) {
	s := qs.Get(key)
	if s == "" {
		return defaultValue, nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (app *application) background(fn func()) {
	app.wg.Add(1)
	go func() {
		defer app.wg.Done()

		// recover panic in background goroutine, instead of terminate the app, log the error
		defer func() {
			if err := recover(); err != nil {
				slog.Error(fmt.Sprintf("%s", err))
			}
		}()

		fn()

	}()
}
