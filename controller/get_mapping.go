package controller

import (
    . "github.com/NamSoGong/DomusPopuli-API/service"
    "github.com/gin-gonic/gin"
)

func getAllHouse(ctx *gin.Context) {
    tok := ctx.Param("token")

    var sid uint
    if err := getSid(tok, &sid); err != nil {
        responseAllHouses(ctx, err, nil); return
    }

    houses, err := GetAllHouse(sid)
    if err != nil {
        responseAllHouses(ctx, err, nil); return
    }

    responseAllHouses(ctx, nil, houses)
}
