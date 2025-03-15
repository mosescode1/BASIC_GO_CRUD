package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
)


var configPath = flag.String("config", ".env", "path to config file")

func main(){
	flag.Parse()

	app := fiber.New()


	log.Fatal(app.Listen(":3000"))
}