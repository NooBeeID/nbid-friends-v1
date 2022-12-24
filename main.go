package main

import (
	"backend/apps"
	"backend/config"
	"backend/pkg/database"
	"os"
)

func main() {
	config.LoadConfig("./cmd/.env")
	sqlDb, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	db := database.NewDatabase()
	db = db.SetSql(sqlDb)

	port := os.Getenv("PORT")
	factory := apps.NewRouterFactory(db)
	router, err := factory.Create(apps.ROUTER_Gin, ":"+port)
	if err != nil {
		panic(err)
	}

	executor := apps.NewRouterExecutor()
	executor.Execute(router)
}
