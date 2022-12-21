package router

import (
	"backend/apps/domain/auth"
	"backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	port   string
	router *gin.Engine
	db     *database.Database
}

func NewRouterGin(port string, db *database.Database) *Gin {
	router := gin.Default()
	return &Gin{
		port:   port,
		router: router,
		db:     db,
	}
}

func (r *Gin) BuildRoutes() {
	v1 := r.router.Group("v1")

	authRoute := auth.NewRouterAuth(v1, r.db)
	authRoute.RegisterAuthRoutes()
}

func (r *Gin) Run() {
	r.router.Run(r.port)
}
