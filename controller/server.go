package controller

import (
    "github.com/gin-gonic/gin"
)

func Serve() {
    r := gin.Default()
    
    r.POST("/signup", signUp)
    r.POST("/signin", signIn)
    r.POST("/preferences", preference)
    r.POST("/house", houseGen)
    r.POST("/checklist/:housepath", writeCheckList)

    r.GET("/house/:token", getAllHouse)

    r.Run(":8080")
}
