package service

import (
    "github.com/NamSoGong/DomusPopuli-API/repository"
    "github.com/NamSoGong/DomusPopuli-API/domain"
)

func SignUp(idpw domain.Sign_t) error {
    if err := repository.CreateUser(idpw.Id, idpw.Pw); err != nil {
        return err
    }
    return nil
}

func SignIn(idpw domain.Sign_t) (int, error) {
    user, err := repository.SelectUser(idpw.Id, idpw.Pw)
    if err != nil {
        return -1, err
    }
    return int(user.ID), nil
}
