package service

import (
    "github.com/NamSoGong/DomusPopuli-API/domain"
)

func CalcEnvScore(coords *domain.Coordinate_t, cgroups []string) (float64, error) {
    if len(cgroups) < 1 {
        return 0, nil
    }

    cnt := 0
    for _, cgroup := range cgroups {
        resp, err := CategorySearch(*coords, cgroup, "500")
        if err != nil {
            return -1, err
        }
        cnt += resp.Meta.TotalCount
    }

    return float64(cnt) / float64(len(cgroups)), nil
}
