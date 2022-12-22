package middleware

import (
	"backend/apps/commons/response"
	"backend/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

type MiddlewareGin struct {
}

func NewMiddlewareGin() *MiddlewareGin {
	return &MiddlewareGin{}
}

func (m *MiddlewareGin) ValidateAuth(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	tokenString := strings.Split(auth, "Bearer ")
	if len(tokenString) != 2 {
		resp := response.UnauthorizedErrorWithAdditionalInfo("len token must be 2")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	tok, err := token.ValidateToken(tokenString[1])

	if err != nil {
		resp := response.UnauthorizedErrorWithAdditionalInfo(err.Error())
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	ctx.Set("authId", tok.AuthId)
	ctx.Next()
}
