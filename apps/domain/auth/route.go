package auth

import (
	"backend/apps/domain/auth/controller"
	"backend/apps/domain/auth/repositories/postgres"
	"backend/apps/domain/auth/services/usecase"
	"backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type RouteAuth interface {
	RegisterAuthRoutes()
}

type routeAuthImpl struct {
	route *gin.RouterGroup
	auth  *controller.ControllerAPI
}

func NewRouterAuth(r *gin.RouterGroup, db *database.Database) RouteAuth {

	repo := postgres.NewAuthRepo(db.DbSQL)
	svc := usecase.NewAuthSvc(repo)
	handler := controller.NewControllerAPI(svc)
	return &routeAuthImpl{
		route: r,
		auth:  handler,
	}
}

func (r *routeAuthImpl) RegisterAuthRoutes() {
	auth := r.route.Group("auth")
	{
		auth.POST("/register", r.auth.Register)
		auth.POST("/login", r.auth.Login)
	}
}
