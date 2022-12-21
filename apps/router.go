package apps

import (
	"backend/apps/router"
	"backend/pkg/database"
	"fmt"
	"log"
)

type Router interface {
	BuildRoutes()
	Run()
}

const (
	ROUTER_Gin = "gin"
)

type RouterFactory struct {
	db *database.Database
}

func NewRouterFactory(db *database.Database) *RouterFactory {

	return &RouterFactory{
		db: db,
	}
}

func (r *RouterFactory) Create(typeRouter string, port string) (Router, error) {
	// fmt.Printf("%+v\n", r.db)

	switch typeRouter {
	case ROUTER_Gin:
		return router.NewRouterGin(port, r.db), nil
	default:
		log.Println()
		return nil, fmt.Errorf("router with type %s not found", typeRouter)
	}
}

type RouterExecutor struct{}

func NewRouterExecutor() *RouterExecutor {
	return &RouterExecutor{}
}

func (*RouterExecutor) Execute(router Router) {
	router.BuildRoutes()
	router.Run()
}
