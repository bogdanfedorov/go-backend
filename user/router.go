package user

import (
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
			ctx.IndentedJSON(http.StatusOK, auth)
		}
	}
}
