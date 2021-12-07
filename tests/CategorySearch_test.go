package tests

import (
	"fmt"
	"testing"

	"github.com/NamSoGong/DomusPopuli-API/domain/coords"
	"github.com/NamSoGong/DomusPopuli-API/service"
)

func TestCategorySearch(t *testing.T) {
    res, err := service.CategorySearch(coords.Coordinate_t{ Long: "127.344608062652", Lat: "36.359011517952"}, "MT1", "20000")

    if err != nil {
        t.Error(err)
        return
    }

    fmt.Println(*res)
}
