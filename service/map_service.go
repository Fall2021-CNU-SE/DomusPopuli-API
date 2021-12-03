package service

import (
	"errors"
	"net/http"
    "net/url"
    "encoding/json"
    "os"

    "github.com/NamSoGong/DomusPopuli-API/domain"
)

func AddressToCoordinate(addr string) (*domain.Coordinate_t, error) {

    // Generate URL
    apiKey := os.Getenv("KAKAO_AK")
    if apiKey == "" {
        return nil, errors.New("Do 'export KAKAO_AK={Kakao REST API Key}'")
    }
    url := `https://dapi.kakao.com/v2/local/search/address.json?query=` + url.QueryEscape(addr)

    // Generate Request
    req, err := http.NewRequest("GET", url, nil)
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
    jsonBody := domain.Response_t{}

    dec.Decode(&jsonBody)

    if jsonBody.Meta.TotalCount < 1 || len(jsonBody.Documents) < 1 {
        return nil, errors.New("Invalid address")
    }

    return &domain.Coordinate_t{
        Long: jsonBody.Documents[0].X,
        Lat: jsonBody.Documents[0].Y,
    }, nil
}
