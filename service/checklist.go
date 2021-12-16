package service

import (
    "github.com/NamSoGong/DomusPopuli-API/domain"
    "github.com/NamSoGong/DomusPopuli-API/repository"
)

func WriteCheckList(sid uint, name string, clist domain.CheckList_t) error {
    user, err := repository.SelectUserBySID(sid)
    if err != nil {
        return err
    }

    score := CalcCheckList(user.PreferedEnv, clist)

    return repository.UpdateCheckList(name, clist, score)
}
