package controller

import (
    "github.com/NamSoGong/DomusPopuli-API/domain"
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
        responseErrorOnly(ctx, err)
    } else {
        responseErrorOnly(ctx, service.SignIn(info))
    }
}

func preference(ctx *gin.Context) {

}

func houseGen(c *gin.Context) {

}

func writeCheckList(c *gin.Context) {

}
