package tests

import (
    "testing"

    "github.com/NamSoGong/DomusPopuli-API/service"
)

func TestCategorySearch(t *testing.T) {
    _, err := service.CategorySearch("127.344608062652", "36.359011517952", "MT1", "20000")
    if err != nil {
        panic(err)
    }
}
