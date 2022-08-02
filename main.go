package main

import (
	_migration "news/databases/migration"
	_config "news/databases/config"
	_factory "news/app"
	_routes "news/routes"
	// _middlewares "news/app/middlewares"
)

func main() {
	// connection database
	dbConn := _config.InitDB()
	// migration table
	_migration.Migration(dbConn)
	// routes
	presenter := _factory.InitFactory(dbConn)
	e := _routes.New(presenter)

	// _middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}