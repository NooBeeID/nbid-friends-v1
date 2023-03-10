package router

import (
	"backend/apps/commons/middleware"
	"backend/apps/domain/auth"
	"backend/apps/domain/follow"
	"backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	port   string
	router *gin.Engine
	db     *database.Database
	middle *middleware.MiddlewareGin
}

func NewRouterGin(port string, db *database.Database) *Gin {
	router := gin.Default()
	midle := middleware.NewMiddlewareGin()
	return &Gin{
		port:   port,
		router: router,
		db:     db,
		middle: midle,
	}
}

func (r *Gin) BuildRoutes() {

	r.router.Use(CORS)

	v1 := r.router.Group("v1")

	authRoute := auth.NewRouterAuth(v1, r.db, r.middle)
	authRoute.RegisterAuthRoutes()

	// use middleware auth
	v1.Use(r.middle.ValidateAuth)
	friendRoute := follow.NewRouterFollow(v1, r.db)
	friendRoute.RegisterFollowRoutes()
}

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}

func (r *Gin) Run() {
	r.router.Run(r.port)
}
