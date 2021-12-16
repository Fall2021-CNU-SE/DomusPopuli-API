package service

import (
    "strings"

    "github.com/NamSoGong/DomusPopuli-API/domain/api"
    "github.com/NamSoGong/DomusPopuli-API/domain/db"
    "github.com/NamSoGong/DomusPopuli-API/repository"
)

func HouseGen(sid uint, houseForm api.MakeHouse_t) error {
    coords, err := AddressToCoordinate(houseForm.Address)
    if err != nil {
        return err
    }

    user, err := repository.SelectUserBySID(sid)
    if err != nil {
        return err
    }

    score, err := CalcEnvScore(coords, strings.Split(user.PreferedFac, "/"))
    if err != nil {
        return err
    }

    if err := repository.CreateHouse(&db.House_t{
        Owner: sid,
        Name: houseForm.Name,
        RentalFee: houseForm.RentalFee,
        Deposit: houseForm.Deposit,
        Maintenance: houseForm.Maintenance,
        RoomSize: houseForm.RoomSize,
        IsRent: houseForm.IsRent,
        Location: *coords,
        EnvScore: score,
        CLScore: 0,
    }); err != nil {
        return err
    }

    return nil
}
