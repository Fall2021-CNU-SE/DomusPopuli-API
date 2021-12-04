package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/NamSoGong/DomusPopuli-API/domain/category"
)

func CategorySearch(long, lat, cgroup, radius string) (*category.Response_t, error) {

    // Get Kakao API key
    apiKey := os.Getenv("KAKAO_AK")
    if apiKey == "" {
        return nil, errors.New("Do 'export KAKAO_AK={Kakao REST API Key}'")
    }

    // Generate URL
    cgroupKey := url.QueryEscape("category_group_code")
    reqUrl := "https://dapi.kakao.com/v2/local/search/category.json"
    reqUrl += fmt.Sprintf("?x=%s&y=%s&%s=%s&radius=%s", long, lat, cgroupKey, cgroup, radius)

    // Generate Request
    req, err := http.NewRequest("GET", reqUrl, nil)
    if err != nil { return nil, err }
    req.Header.Add("Authorization", "KakaoAK " + apiKey)

    // Send Request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil { return nil, err }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return nil, errors.New("Request responed with " + resp.Status)
    }

    // Parse JSON
    dec := json.NewDecoder(resp.Body)
    jsonBody := category.Response_t{}

    dec.Decode(&jsonBody)

    if jsonBody.Meta.TotalCount < 1 || len(jsonBody.Documents) < 1 {
        return nil, errors.New("Invalid address")
    }

    fmt.Println(jsonBody)

    return nil, nil
}
