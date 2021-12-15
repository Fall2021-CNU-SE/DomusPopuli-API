package repository

import (
    "errors"

    "github.com/NamSoGong/DomusPopuli-API/exceptions"
    . "github.com/NamSoGong/DomusPopuli-API/domain/db"
    "gorm.io/gorm"
)

func CreateUser(id, pw string) error {
    db, err := getDB()
    if err != nil {
        return err
    }

    if res := db.Create(&User_t{
        SignID: id,
        PW: pw,
    }); res.Error != nil {
        return res.Error
    }

    return nil
}

func SelectUser(id, pw string) (*User_t, error) {
    db, err := getDB()
    if err != nil {
        return nil, err
    }

    var user User_t
    res := db.Where("sign_id = ? AND pw = ?", id, pw).First(&user)

    if errors.Is(res.Error, gorm.ErrRecordNotFound) {
        return nil, exceptions.UserNotFound
    } else if res.Error != nil {
        return nil, res.Error
    }

    return &user, nil
}
