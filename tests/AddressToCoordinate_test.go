package tests

import (
    "testing"

    "github.com/NamSoGong/DomusPopuli-API/service"
)

func TestAddressToCoordinate(t *testing.T) {
    coor, err := service.AddressToCoordinate("문화원로 80")
    if err != nil {
        t.Error(err)
    }

    t.Log("Longitude " + coor.Long + " Latitude " + coor.Lat)
}
