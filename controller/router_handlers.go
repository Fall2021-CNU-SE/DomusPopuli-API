package controller

import (
	"github.com/NamSoGong/DomusPopuli-API/domain"
	"github.com/NamSoGong/DomusPopuli-API/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
    err := ctx.ShouldBindJSON(&info)
    if err != nil {
        responseErrorOnly(ctx, err); return
    }

    sid, err := service.SignIn(info)
    if err != nil {
        responseErrorOnly(ctx, err); return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sid": sid,
    })
    
    tokenStr, err := token.SignedString([]byte(TestSecrkey))
    if err != nil {
        responseErrorOnly(ctx, err); return
    }

    ctx.SetCookie("sid", tokenStr, CookieExp, "/", Domain, false, false)
    responseErrorOnly(ctx, nil)
}

func preference(ctx *gin.Context) {

}

func houseGen(c *gin.Context) {

}

func writeCheckList(c *gin.Context) {

}
