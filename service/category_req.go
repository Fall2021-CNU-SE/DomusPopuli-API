package service

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/url"

    "github.com/NamSoGong/DomusPopuli-API/domain/category"
    "github.com/NamSoGong/DomusPopuli-API/domain/coords"
)

func CategorySearch(coor coords.Coordinate_t, cgroup, radius string) (*category.Response_t, error) {
    // Generate URL
    cgroupKey := url.QueryEscape("category_group_code")
    reqUrl := "https://dapi.kakao.com/v2/local/search/category.json"
    reqUrl += fmt.Sprintf("?x=%s&y=%s&%s=%s&radius=%s", coor.Long, coor.Lat, cgroupKey, cgroup, radius)

    bodyReader, err := KakaoMAPIGet(reqUrl)
    if err != nil {
        return nil, err
    }

    // Parse JSON
    dec := json.NewDecoder(bodyReader)
    jsonBody := category.Response_t{}

    dec.Decode(&jsonBody)

    if jsonBody.Meta.TotalCount < 1 || len(jsonBody.Documents) < 1 {
        return nil, errors.New("Invalid address")
    }

    return &jsonBody, nil
}
