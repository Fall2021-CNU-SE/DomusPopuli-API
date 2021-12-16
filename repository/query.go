package repository

import (
    "errors"

    "github.com/NamSoGong/DomusPopuli-API/domain"
    . "github.com/NamSoGong/DomusPopuli-API/domain/db"
    "github.com/NamSoGong/DomusPopuli-API/exceptions"
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

func UpdatePreferences(sid, budget uint, workAddr *domain.Coordinate_t, fac string, env domain.CheckList_t) error {
    db, err := getDB()
    if err != nil {
        return err
    }

    if res := db.Model(&User_t{}).Where("id = ?", sid).Updates(User_t{
        Budget: budget,
        WorkAddress: *workAddr,
        PreferedFac: fac,
        PreferedEnv: env,
    }); res.Error != nil {
        return res.Error
    }

    return nil
}

func CreateHouse(house *House_t) error {
    db, err := getDB()
    if err != nil {
        return err
    }

    if res := db.Create(house); res.Error != nil {
        return res.Error
    }

    return nil
}
