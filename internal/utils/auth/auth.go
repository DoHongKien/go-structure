package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(ctx *gin.Context) (string, bool) {

	token := ctx.Request.Header.Get("Authorization")

	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer "), true
	}

	return "", false
}