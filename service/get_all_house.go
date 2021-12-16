package service

import (
    "github.com/NamSoGong/DomusPopuli-API/repository"
    "github.com/NamSoGong/DomusPopuli-API/domain/api"
)

func GetAllHouse(sid uint) ([]api.HouseSummary_t, error) {
    houses, err := repository.SelectAllHouse(sid)
    if err != nil {
        return nil, err
    }

    summaries := make([]api.HouseSummary_t, len(houses))
    for i := range houses {
        summaries[i].Name = houses[i].Name
        summaries[i].TotalScore = houses[i].CLScore * houses[i].EnvScore
    }

    return summaries, nil
}
