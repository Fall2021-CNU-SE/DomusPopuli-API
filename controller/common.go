package controller

import (
    "errors"
    "net/http"

    "github.com/NamSoGong/DomusPopuli-API/exceptions"
    "github.com/gin-gonic/gin"
)

func httpErrorCode(err error) int {
    if errors.Is(err, exceptions.UserNotFound) {
        return http.StatusUnauthorized
    }

    return http.StatusBadRequest
}

func responseErrorOnly(ctx *gin.Context, err error) {
    if err != nil {
        ctx.JSON(httpErrorCode(err), gin.H{ "error": err.Error(), })
    } else {
        ctx.JSON(http.StatusOK, gin.H{ "error": nil })
    }
}
