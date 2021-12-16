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
        PreferedEnv: domain.CheckList_t{
            H0: 1, H1: 1, H2: 1,
            B0: 1, B1: 1, B2: 1,
            K0: 1, K1: 1, K2: 1,
            O0: 1, O1: 1, O2: 1,
            E0: 1,
        },
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

func UpdateCheckList(name string, clist domain.CheckList_t) error {
    db, err := getDB()
    if err != nil {
        return err
    }

    if res := db.Model(&House_t{}).Where("name = ?", name).Updates(House_t{
        CheckList: clist,
    }); res.Error != nil {
        return res.Error
    }

    return nil
}
