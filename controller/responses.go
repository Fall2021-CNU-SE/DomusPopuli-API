package controller

import (
    "errors"
    "net/http"

    "github.com/NamSoGong/DomusPopuli-API/exceptions"
    "github.com/NamSoGong/DomusPopuli-API/domain/api"
    "github.com/gin-gonic/gin"
)

func httpStatusCode(err error) int {
    switch {
    case err == nil:                                return http.StatusOK
    case errors.Is(err, exceptions.UserNotFound):   return http.StatusUnauthorized
    default:                                        return http.StatusBadRequest
    }
}

func errMsg(err error) string {
    switch {
    case err != nil:    return err.Error()
    default:            return "null"
    }
}

func responseErrorOnly(ctx *gin.Context, err error) {
    ctx.JSON(httpStatusCode(err), gin.H{
        "error": errMsg(err),
    })
}

func responseSignIn(ctx *gin.Context, err error, token string) {
    ctx.JSON(httpStatusCode(err), gin.H{
        "error": errMsg(err),
        "token": token,
    })
}

func responseAllHouses(ctx *gin.Context, err error, houses []api.HouseSummary_t) {
    ctx.JSON(httpStatusCode(err), gin.H{
        "error": errMsg(err),
        "houses": houses,
    })
}
