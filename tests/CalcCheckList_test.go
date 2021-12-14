package tests

import (
    "testing"
    "math"

    "github.com/NamSoGong/DomusPopuli-API/service"
    "github.com/NamSoGong/DomusPopuli-API/domain"
)

func TestCalcCheckList(t *testing.T) {
    res := service.CalcCheckList(domain.CheckList_t{
        H0: 2, H1: 1, H2: 2,
        B0: 1, B1: 2, B2: 1,
        K0: 2, K1: 1, K2: 2,
        O0: 1, O1: 2, O2: 1,
        E0: 2,
    }, domain.CheckList_t{
        H0: 1, H1: 2, H2: 3,
        B0: 4, B1: 5, B2: 1,
        K0: 2, K1: 3, K2: 4,
        O0: 5, O1: 1, O2: 2,
        E0: 3,
    })

    if math.Abs(4.23076923 - res) < 0.001 {
        t.Error("FAIL")
    }
}
