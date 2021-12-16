package controller

import (
	"github.com/NamSoGong/DomusPopuli-API/domain"
	"github.com/NamSoGong/DomusPopuli-API/domain/api"
	"github.com/NamSoGong/DomusPopuli-API/service"
	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
    var info domain.Sign_t
    if err := ctx.ShouldBindJSON(&info); err != nil {
        responseErrorOnly(ctx, err)
    } else {
        responseErrorOnly(ctx, service.SignUp(info))
    }
}

func signIn(ctx *gin.Context) {
    var info domain.Sign_t
    if err := ctx.ShouldBindJSON(&info); err != nil {
        responseSignIn(ctx, err, ""); return
    }

    sid, err := service.SignIn(info)
    if err != nil {
        responseSignIn(ctx, err, ""); return
    }

    tok, err := signSid(sid)
    if err != nil {
        responseSignIn(ctx, err, ""); return
    }
    
    responseSignIn(ctx, err, tok)
}

func preference(ctx *gin.Context) {
    var pref api.Preference_t
    if err := ctx.ShouldBindJSON(&pref); err != nil {
        responseErrorOnly(ctx, err); return
    }

    var sid uint
    if err := getSid(pref.Token, &sid); err != nil {
        responseErrorOnly(ctx, err); return
    }

    if err := service.UpdatePreferences(sid, pref); err != nil {
        responseErrorOnly(ctx, err); return
    }

    responseErrorOnly(ctx, nil)
}

func houseGen(ctx *gin.Context) {
    var house api.MakeHouse_t
    if err := ctx.ShouldBindJSON(&house); err != nil {
        responseErrorOnly(ctx, err); return
    }

    var sid uint
    if err := getSid(house.Token, &sid); err != nil {
        responseErrorOnly(ctx, err); return
    }

    if err := service.HouseGen(sid, house); err != nil {
        responseErrorOnly(ctx, err); return
    }

    responseErrorOnly(ctx, nil)
}

func writeCheckList(ctx *gin.Context) {
    var checklist api.CheckList_t
    if err := ctx.ShouldBindJSON(&checklist); err != nil {
        responseErrorOnly(ctx, err); return
    }

    houseName := ctx.Param("housepath")

    if err := service.WriteCheckList(houseName, checklist.Checklist); err != nil {
        responseErrorOnly(ctx, err); return
    }

    responseErrorOnly(ctx, nil)
}
