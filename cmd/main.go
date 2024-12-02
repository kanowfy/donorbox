package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	appKey := os.Getenv("DROPBOX_APPKEY")
	appSecret := os.Getenv("DROPBOX_APPSECRET")
	accessCode := os.Getenv("ACCESS_CODE")
	req, err := http.NewRequest("POST", "https://api.dropbox.com/oauth2/token", nil)
	if err != nil {
		log.Fatal(err)
	}

	base64Auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", appKey, appSecret)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64Auth))

	qs := req.URL.Query()
	qs.Add("code", accessCode)
	qs.Add("grant_type", "authorization_code")

	req.URL.RawQuery = qs.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %s", body)
}
