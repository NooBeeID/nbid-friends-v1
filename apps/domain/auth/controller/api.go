package controller

import (
	"backend/apps/commons/response"
	"backend/apps/domain/auth/params"
	"backend/apps/domain/auth/services"

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

func (c *ControllerAPI) Search(ctx *gin.Context) {
	var req = new(params.UserSearchRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		resp := response.GeneralError()
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	authId := ctx.GetInt("authId")
	if authId == 0 {
		resp := response.GeneralErrorWithAdditionalInfo("auth id cannot be 0")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	req.AuthId = authId

	result, custErr := c.svc.Search(ctx.Request.Context(), req.Email, req.AuthId)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("SEARCH SUCCESS", result)
	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) Profile(ctx *gin.Context) {
	authId := ctx.GetInt("authId")
	if authId == 0 {
		resp := response.GeneralErrorWithAdditionalInfo("auth id cannot be 0")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	profile, custErr := c.svc.Profile(ctx, authId)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("GET PROFILE SUCCESS", profile)
	ctx.JSON(resp.StatusCode, resp)
}
