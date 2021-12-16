package controller

import (
    . "github.com/NamSoGong/DomusPopuli-API/service"
    "github.com/gin-gonic/gin"
)

func getAllHouse(ctx *gin.Context) {
    var sid uint
    if err := getSid(ctx.Param("token"), &sid); err != nil {
        responseAllHouses(ctx, err, nil); return
    }

    houses, err := GetAllHouse(sid)
    if err != nil {
        responseAllHouses(ctx, err, nil); return
    }

    responseAllHouses(ctx, nil, houses)
}

func getHouseDetail(ctx *gin.Context) {
    var sid uint
    if err := getSid(ctx.Param("token"), &sid); err != nil {
        responseHouseDetail(ctx, err, nil)
    }

    detail, err := GetHouseDetail(sid, ctx.Param("housepath"))
    responseHouseDetail(ctx, err, detail)
}
