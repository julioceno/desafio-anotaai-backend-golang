package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetValueByParams(ctx *gin.Context, value string) (string, error) {
	id := ctx.Param(value)
	if strings.TrimSpace(id) == "" {
		return "", errors.New(fmt.Sprintf("%s param not exist", value))
	}

	return id, nil
}
