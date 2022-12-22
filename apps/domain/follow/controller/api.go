package controller

import (
	"backend/apps/commons/response"
	"backend/apps/domain/follow/params"
	"backend/apps/domain/follow/services"

	"github.com/gin-gonic/gin"
)

type ControllerAPI struct {
	svc services.FollowSvc
}

func NewControllerAPI(svc services.FollowSvc) *ControllerAPI {
	return &ControllerAPI{
		svc: svc,
	}
}

func (c *ControllerAPI) Follow(ctx *gin.Context) {
	var req = params.FollowingRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		resp := response.GeneralErrorWithAdditionalInfo(err.Error())
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

	custErr := c.svc.FollowFriend(ctx, &req)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.CreatedSuccess()
	ctx.JSON(resp.StatusCode, resp)
}

func (c *ControllerAPI) GetMyFollowing(ctx *gin.Context) {
	authId := ctx.GetInt("authId")

	if authId == 0 {
		resp := response.GeneralErrorWithAdditionalInfo("auth id cannot be 0")
		ctx.AbortWithStatusJSON(resp.StatusCode, resp)
		return
	}

	myFollowing, custErr := c.svc.GetMyFollowing(ctx, authId)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	resp := response.GeneralSuccessCustomMessageAndPayload("GET ALL SUCCESS", myFollowing)
	ctx.JSON(resp.StatusCode, resp)
}
