package helper

import (
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
)

func ReadString(qs url.Values, key, defaultValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	return s
}

func ReadInt(qs url.Values, key string, defaultValue int) (int, error) {
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

func Background(fn func()) {
	go func() {
		// recover panic in background goroutine, instead of terminate the app, log the error
		defer func() {
			if err := recover(); err != nil {
				slog.Error(fmt.Sprintf("%s", err))
			}
		}()

		fn()
	}()
}
