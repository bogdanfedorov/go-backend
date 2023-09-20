package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func newUser(ctx *gin.Context) {
	var body User
	if err := ctx.BindJSON(&body); err == nil {
		if user, createUserError := createNewUser(body); createUserError != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			ctx.IndentedJSON(http.StatusCreated, user)
		}
	}
}

func loginUser(ctx *gin.Context) {
	var body User
	if err := ctx.BindJSON(&body); err == nil {
		if auth, loginError := login(body); loginError != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			session := sessions.Default(ctx)
			session.Set("UserId", auth.UserID)

			if err := session.Save(); err != nil {
				ctx.AbortWithStatus(http.StatusBadRequest)
				return
			}
			ctx.IndentedJSON(http.StatusOK, nil)
		}
	}
}

func Bind(router *gin.Engine) {
	userRouter := router.Group("/users")

	userRouter.POST("/", newUser)
	userRouter.POST("/login", loginUser)
}
