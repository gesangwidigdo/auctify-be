package utils

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"status":  "success",
		"data":    data,
	})
}

func FailResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "fail",
		"message": message,
	})
}