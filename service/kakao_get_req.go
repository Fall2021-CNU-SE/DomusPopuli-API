package service

import (
    "errors"
    "net/http"
    "os"
    "io"
)

func KakaoMAPIGet(reqUrl string) (io.Reader, error) {

    // Get Kakao API key
    apiKey := os.Getenv("KAKAO_AK")
    if apiKey == "" {
        return nil, errors.New("Do 'export KAKAO_AK={Kakao REST API Key}'")
    }

    // Generate Request
    req, err := http.NewRequest("GET", reqUrl, nil)
    if err != nil { return nil, err }
    req.Header.Add("Authorization", "KakaoAK " + apiKey)

    // Send Request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil { return nil, err }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("Request responed with " + resp.Status)
    }

    return resp.Body, nil
}
