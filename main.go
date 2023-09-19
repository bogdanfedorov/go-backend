package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-backend/user"
)

func main() {
	router := gin.Default()
	user.Bind(router)

	if err := router.Run(":3000"); err != nil {
		fmt.Println(err.Error())
		return
	}
}
