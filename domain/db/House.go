package db

import (
    "gorm.io/gorm"

    . "github.com/NamSoGong/DomusPopuli-API/domain"
)

type House_t struct {
    gorm.Model

    // Name of the House
    Name        string

    // Monthly Retal Fee
    RentalFee   uint

    // Cash Deposited
    Deposit     uint

    // Location of the House
    Location    Coordinate_t    `gorm:"embedded"`

    // Scores for Check List Entry
    CheckList   CheckList_t     `gorm:"embedded"`
}
