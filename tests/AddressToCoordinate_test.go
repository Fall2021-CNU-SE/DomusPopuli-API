package tests

import (
    "testing"
    "fmt"

    "github.com/NamSoGong/DomusPopuli-API/service"
)

func TestAddressToCoordinate(t *testing.T) {
    coords, err := service.AddressToCoordinate("문화원로 80")
    if err != nil {
        t.Error(err)
        return
    }

    fmt.Println(coords.Long)
    fmt.Println(coords.Lat)
}
