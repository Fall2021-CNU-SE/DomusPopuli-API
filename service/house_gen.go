package service

import (
    "github.com/NamSoGong/DomusPopuli-API/domain/db"
    "github.com/NamSoGong/DomusPopuli-API/domain/api"
    "github.com/NamSoGong/DomusPopuli-API/repository"
)

func HouseGen(sid uint, houseForm api.MakeHouse_t) error {
    coords, err := AddressToCoordinate(houseForm.Address1)
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
        Address2: houseForm.Address2,
        Location: *coords,
    }); err != nil {
        return err
    }

    return nil
}
