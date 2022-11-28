package api

import (
	"github.com/fahturr/default_project/internal/app"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, app *app.YourApp) {
	mainRoute := r.Group("/v1")
	{
		mainRoute.GET("/")
	}
}
