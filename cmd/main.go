package main

import (
	"github.com/fahturr/default_project/api"
	"github.com/fahturr/default_project/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()

	yourApp := app.NewYourApp()

	api.RegisterRoutes(router, yourApp)

	err := router.Run(":9292")
	if err != nil {
		panic(err)
	}
}
