package service

import (
    "reflect"

    "github.com/NamSoGong/DomusPopuli-API/domain"
)

func CalcCheckList(prefs, scores domain.CheckList_t) float64 {
    refPref := reflect.ValueOf(prefs)
    refScore := reflect.ValueOf(scores)

    var acc float64 = 0
    for i := 0; i < refPref.NumField(); i++ {
        pref := refPref.Field(i).Interface().(uint)
        score := refScore.Field(i).Interface().(uint)
        acc += float64(pref * score) / 3
    }

    return acc / float64(refPref.NumField())
}

func PrefWeight(prefs []string) *domain.CheckList_t {
    clist := &domain.CheckList_t{}

    for _, pref := range prefs {
        switch pref {
        case "H0": clist.H0 = 2
        case "H1": clist.H1 = 2
        case "H2": clist.H2 = 2
        case "B0": clist.B0 = 2
        case "B1": clist.B1 = 2
        case "B2": clist.B2 = 2
        case "K0": clist.K0 = 2
        case "K1": clist.K1 = 2
        case "K2": clist.K2 = 2
        case "O0": clist.O0 = 2
        case "O1": clist.O1 = 2
        case "O2": clist.O2 = 2
        case "E0": clist.E0 = 2
        }
    }

    return clist
}
