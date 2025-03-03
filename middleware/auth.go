package middleware

import (
	"net/http"
	"time"

	"github.com/gesangwidigdo/auctify-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		utils.FailResponse(c, http.StatusUnauthorized, "Unauthorized")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := utils.VerifyToken(tokenString)
	if err != nil {
		utils.FailResponse(c, http.StatusUnauthorized, "Token verification failed")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.FailResponse(c, http.StatusUnauthorized, "Token expired")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("username", claims["sub"])
		c.Set("id", claims["id"])
		c.Set("role", claims["role"])
		c.Next()
	} else {
		utils.FailResponse(c, http.StatusUnauthorized, "Token verification failed")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
