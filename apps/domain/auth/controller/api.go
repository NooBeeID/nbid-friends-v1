package controller

import (
	"backend/apps/domain/auth/params"
	"backend/apps/domain/auth/services"
	"backend/helper/response"

	"github.com/gin-gonic/gin"
)

type ControllerAPI struct {
	svc services.AuthSvc
}

func NewControllerAPI(svc services.AuthSvc) *ControllerAPI {
	return &ControllerAPI{
		svc: svc,
	}
}

func (c *ControllerAPI) Register(ctx *gin.Context) {
	var req = new(params.UserRegisterRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		resp := response.GeneralError()
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	custErr := c.svc.Register(ctx.Request.Context(), req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccess()
	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) Login(ctx *gin.Context) {
	var req = new(params.UserLoginRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		resp := response.GeneralError()
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	result, custErr := c.svc.Login(ctx.Request.Context(), req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("LOGIN SUCCESS", result)
	ctx.JSON(resp.StatusCode, resp)

}
