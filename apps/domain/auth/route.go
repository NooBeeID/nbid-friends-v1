package auth

import (
	"backend/apps/commons/middleware"
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
	route  *gin.RouterGroup
	auth   *controller.ControllerAPI
	middle *middleware.MiddlewareGin
}

func NewRouterAuth(r *gin.RouterGroup, db *database.Database, middle *middleware.MiddlewareGin) RouteAuth {

	repo := postgres.NewAuthRepo(db.DbSQL)
	svc := usecase.NewAuthSvc(repo)
	handler := controller.NewControllerAPI(svc)
	return &routeAuthImpl{
		route:  r,
		auth:   handler,
		middle: middle,
	}
}

func (r *routeAuthImpl) RegisterAuthRoutes() {
	auth := r.route.Group("auth")
	{
		auth.POST("/register", r.auth.Register)
		auth.POST("/login", r.auth.Login)
		auth.POST("/search", r.auth.Search)

		auth.Use(r.middle.ValidateAuth)
		auth.GET("/profile", r.auth.Profile)
	}
}
