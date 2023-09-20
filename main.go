package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"go-backend/user"
)

type Config struct {
	Second int
	Minute int
	Hour   int
	Day    int

	Port string
}

func GetConfig() Config {
	config := Config{}

	config.Second = 1
	config.Minute = config.Second * 60
	config.Hour = config.Minute * 60
	config.Day = config.Hour * 24
	config.Port = ":3000"

	return config
}

func main() {
	config := GetConfig()

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge: config.Day})
	router.Use(sessions.Sessions("session", store))

	user.Bind(router)

	if err := router.Run(config.Port); err != nil {
		fmt.Println(err.Error())
		return
	}
}
