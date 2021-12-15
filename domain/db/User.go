package db

import (
    "gorm.io/gorm"

    . "github.com/NamSoGong/DomusPopuli-API/domain"
)

type User_t struct {
    gorm.Model
    SignID          string
    PW              string
    Budget          uint
    WorkAddress     Coordinate_t    `gorm:"embedded"`
    PreferedFac     string
    PreferedEnv     CheckList_t     `gorm:"embedded"`
}
