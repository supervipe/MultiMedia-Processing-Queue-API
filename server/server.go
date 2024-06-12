package main

import (
	"POC-Load-Videos/routes"
	"POC-Load-Videos/tasks"
	"POC-Load-Videos/worker"
	"github.com/gofiber/fiber/v2"
	"log"
)

const redisAddress = "127.0.0.1:6379"

func main() {
	app := fiber.New()
	routes.Setup(app)
	tasks.Init(redisAddress)
	defer tasks.Close()
	go worker.RunWorkers()
	log.Fatal(app.Listen(":3000"))
}
