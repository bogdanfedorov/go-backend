package user

import (
	"github.com/gin-gonic/gin"
)

func Bind(router *gin.Engine) {
	userRouter := router.Group("/user")

	userRouter.POST("/", newUser)
	userRouter.POST("/login", loginUser)
}
