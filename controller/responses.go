package controller

import (
    "errors"
    "net/http"

    "github.com/NamSoGong/DomusPopuli-API/exceptions"
    "github.com/gin-gonic/gin"
)

func httpStatusCode(err error) int {
    if err == nil {
        return http.StatusOK
    }

    if errors.Is(err, exceptions.UserNotFound) {
        return http.StatusUnauthorized
    }

    return http.StatusBadRequest
}

func responseErrorOnly(ctx *gin.Context, err error) {
    if err != nil {
        ctx.JSON(httpStatusCode(err), gin.H{ "error": err.Error(), })
    } else {
        ctx.JSON(httpStatusCode(nil), gin.H{ "error": nil })
    }
}
