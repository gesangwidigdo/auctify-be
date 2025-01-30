package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ExtractID(ctx *gin.Context) (uint, int, error) {
	id, ok := ctx.Get("id")
	if !ok {
		return 0, 404, errors.New("User not found")
	}

	idFloat, ok := id.(float64)
	if !ok {
		return 0, 400, errors.New("Invalid user ID type")
	}

	idUint := uint(idFloat)
	return idUint, 0, nil
}
