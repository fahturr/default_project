package middleware

import (
	"github.com/fahturr/default_project/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"message": "unauthorized access"},
			)
			return
		}

		token := strings.Split(bearerToken, "Bearer ")
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"message": "token is empty"},
			)
			return
		}

		claims, err := helper.ValidateToken(token[1])
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"message": "token invalid"},
			)
			return
		}

		userId := claims["uid"].(float64)
		outletID := claims["oid"].(float64)
		userName := claims["uname"].(string)
		outletName := claims["oname"].(string)

		ctx.Set("uid", userId)
		ctx.Set("oid", outletID)
		ctx.Set("uname", userName)
		ctx.Set("oname", outletName)

		ctx.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache_Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}
