package main

import (
	_factory "news/app"
	_middlewares "news/app/middlewares"
	_config "news/databases/config"
	_migration "news/databases/migration"
	_routes "news/routes"
)

func main() {
	// connection database
	dbConn := _config.InitDB()
	// migration table
	_migration.Migration(dbConn)
	// routes
	presenter := _factory.InitFactory(dbConn)
	e := _routes.New(presenter)

	_middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
