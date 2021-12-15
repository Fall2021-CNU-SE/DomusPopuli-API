package controller

import (
	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()
    
    r.POST("/signup", signUp)
    r.POST("/signin", signIn)
    r.POST("/preference", preference)
    r.POST("/house", houseGen)
    r.POST("/checklist", writeCheckList)

    r.Run(":8080")
}
