package db

import (
    "gorm.io/gorm"

    . "github.com/NamSoGong/DomusPopuli-API/domain"
)

type User_t struct {
    gorm.Model
    ID              string
    PW              string
    WorkAddress     Coordinate_t
    PreferedFac     []string
    PreferedEnv     CheckList_t
}
