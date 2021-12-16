package service

import (
    "strings"

    "github.com/NamSoGong/DomusPopuli-API/domain/api"
    "github.com/NamSoGong/DomusPopuli-API/repository"
)

func GetHouseDetail(sid uint, houseName string) (*api.HouseDetail_t, error) {
    house, err := repository.SelectHouse(sid, houseName)
    if err != nil {
        return nil, err
    }

    houseDetail := &api.HouseDetail_t{}
    houseDetail.TotalScore = CalcTotalScore(house.EnvScore, house.CLScore)
    houseDetail.ChecklistScore = house.CLScore
    houseDetail.FacilityScore = house.EnvScore

    user, err := repository.SelectUserBySID(sid)
    if err != nil {
        return nil, err
    }

    facs := make([]api.FacilitySummary_t, 0)
    for _, cgroup := range strings.Split(user.PreferedFac, "/") {
        resp, err := CategorySearch(house.Location, cgroup, "500")
        if err != nil {
            return nil, err
        }

        for _, doc := range resp.Documents {
            facs = append(facs, api.FacilitySummary_t{
                Name: doc.PlaceName,
                Address: doc.RoadAddressName,
            })
        }
    }
    houseDetail.Facilities = facs

    return houseDetail, nil
}
