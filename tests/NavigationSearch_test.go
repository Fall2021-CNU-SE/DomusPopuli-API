package tests

import (
    "fmt"
    "testing"

    "github.com/NamSoGong/DomusPopuli-API/domain"
    "github.com/NamSoGong/DomusPopuli-API/service"
)

func TestNavigationSearch(t *testing.T) {
    res, err := service.NavigationSearch(
        domain.Coordinate_t{
            Long: "127.344608062652",
            Lat: "36.359011517952",
        },
        domain.Coordinate_t{
            Long: "127.338778358001",
            Lat: "36.3589659179602",
        })

    if err != nil {
        t.Error(err)
        return
    }

    fmt.Println(*res)
}
