package follow

import (
	"backend/apps/domain/follow/controller"
	"backend/apps/domain/follow/repositories/postgres"
	"backend/apps/domain/follow/services/usecase"
	"backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type RouteFollow interface {
	RegisterFollowRoutes()
}

type routeFollowImpl struct {
	router *gin.RouterGroup
	follow *controller.ControllerAPI
}

// RegisterFollowRoutes implements RouteFollow
func (r *routeFollowImpl) RegisterFollowRoutes() {

	follow := r.router.Group("follow")
	{
		follow.POST("", r.follow.Follow)
		follow.GET("", r.follow.GetMyFollowing)
		follow.DELETE("/", r.follow.UnfollFriend)
	}
}

func NewRouterFollow(r *gin.RouterGroup, db *database.Database) RouteFollow {
	repo := postgres.NewFollowRepo(db.DbSQL)
	svc := usecase.NewFollowSvc(repo)
	handler := controller.NewControllerAPI(svc)
	return &routeFollowImpl{
		router: r,
		follow: handler,
	}
}
