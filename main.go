package main

import (
	"github.com/gesangwidigdo/auctify-be/config"
	"github.com/gesangwidigdo/auctify-be/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		utils.SuccessResponse(c, 200, "Hello, World!")
	});

	r.Run()
}
