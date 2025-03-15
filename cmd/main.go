package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/config"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/database"
	"github.com/mosescode1/BASIC_GO_CRUD.git/internals/routes"
)


var configPath = flag.String("config", ".env", "path to config file")

func main(){
	flag.Parse()

	appConfig, err := config.LoadConfigFromfile(*configPath)

	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	_, err = database.ConnectDatabase(appConfig.Database.DatabaseName)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = database.MigrateDtabase()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	defer database.CloseConnection()


	// HTTP SERVER STARTS HERE
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": appConfig.Server.AppName + " is running",
		})
	})

	routes.SetUpTodoRoutes(app)

	log.Fatal(app.Listen(":" + appConfig.Server.Port))
}