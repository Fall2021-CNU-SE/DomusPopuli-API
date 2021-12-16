package tests

import (
	"testing"

	"github.com/NamSoGong/DomusPopuli-API/service"
)

func TestPrefWeight(t *testing.T) {
    weight := service.PrefWeight([]string{"H0", "B1", "K2"})

    if weight.H0 == 2 && weight.H1 == 1 && weight.H2 == 1 &&
        weight.B0 == 1 && weight.B1 == 2 && weight.B2 == 1 &&
        weight.K0 == 1 && weight.K1 == 1 && weight.K2 == 2 &&
        weight.O0 == 1 && weight.O1 == 1 && weight.O2 == 1 &&
        weight.E0 == 1 {
        return
    } else {
        t.Error("FAIL")
    }
}
