package service

import (
    "github.com/NamSoGong/DomusPopuli-API/domain"
    "github.com/NamSoGong/DomusPopuli-API/repository"
)

func WriteCheckList(name string, clist domain.CheckList_t) error {
    return repository.UpdateCheckList(name, clist)
}
