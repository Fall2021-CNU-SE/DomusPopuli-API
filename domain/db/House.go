package db

import (
    "gorm.io/gorm"

    . "github.com/NamSoGong/DomusPopuli-API/domain"
)

type House_t struct {
    gorm.Model
    Owner       uint
    Name        string
    RentalFee   uint
    Deposit     uint
    Maintenance uint
    RoomSize    uint
    IsRent      bool
    Location    Coordinate_t    `gorm:"embedded"`
    Address2    string
    CheckList   CheckList_t     `gorm:"embedded"`
    EnvScore    float64
    CLScore     float64
}
