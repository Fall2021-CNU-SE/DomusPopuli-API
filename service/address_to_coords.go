package service

import (
    "encoding/json"
    "errors"
    "net/url"

    "github.com/NamSoGong/DomusPopuli-API/domain/coords"
)

func AddressToCoordinate(addr string) (*coords.Coordinate_t, error) {
    reqUrl := `https://dapi.kakao.com/v2/local/search/address.json?query=` + url.QueryEscape(addr)

    bodyReader, err := KakaoMAPIGet(reqUrl)
    if err != nil {
        return  nil, err
    }
    defer bodyReader.Close()

    // Parse JSON
    dec := json.NewDecoder(bodyReader)
    jsonBody := coords.Response_t{}

    dec.Decode(&jsonBody)

    if jsonBody.Meta.TotalCount < 1 || len(jsonBody.Documents) < 1 {
        return nil, errors.New("Invalid Address")
    }

    return &coords.Coordinate_t{
        Long: jsonBody.Documents[0].X,
        Lat: jsonBody.Documents[0].Y,
    }, nil
}
