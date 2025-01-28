package main

import (
	"github.com/gesangwidigdo/auctify-be/config"
	"github.com/gesangwidigdo/auctify-be/model"
	"github.com/gesangwidigdo/auctify-be/router"
	"github.com/gesangwidigdo/auctify-be/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		utils.SuccessResponse(c, 200, "Hello, World!")
	})

	db := config.ConnectDB()

	if err := model.Migrate(db); err != nil {
		panic(err.Error())
	}

	router.Router(r, db)

	r.Run()
}
