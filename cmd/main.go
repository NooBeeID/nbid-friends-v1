package main

import (
	"backend/apps"
	"backend/pkg/database"
)

func main() {

	sqlDb, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	db := database.NewDatabase()
	db = db.SetSql(sqlDb)

	port := ":4444"
	factory := apps.NewRouterFactory(db)
	router, err := factory.Create(apps.ROUTER_Gin, port)
	if err != nil {
		panic(err)
	}

	executor := apps.NewRouterExecutor()
	executor.Execute(router)
}
