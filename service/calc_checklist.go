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
        acc += float64(pref * score) / 2.5
    }

    return acc / float64(refPref.NumField())
}
