package api

import (
    "github.com/NamSoGong/DomusPopuli-API/domain"
)

type CheckList_t struct {
    Token       string              `json:"token" binding:"required"`
    Checklist   domain.CheckList_t  `json:"checklist"`
}
