package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"team1.asia/fibo/config"
	"team1.asia/fibo/db"
	_ "team1.asia/fibo/docs"
	"team1.asia/fibo/log"
	"team1.asia/fibo/router"
)

// @title Fibo App
// @version 1.0
// @description This is an API for Fibo Application

// @contact.name Toan Dinh
// @contact.email toandptech@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	// Read configuration from the environment variables
	config.LoadConfigFromEnv()

	// Zap logger
	log.SetupLogger()

	// DB Connection
	db.Connect()

	migrate := flag.Bool("migrate", false, "Check the migration request")
	flag.Parse()

	if *migrate {
		// DB Migration
		db.Migrate()
	} else {
		// Create new Fiber application
		app := fiber.New(fiber.Config{
			Concurrency:           256 * 1024 * 1024,
			ServerHeader:          config.App.Server.Name,
			BodyLimit:             config.App.Server.UploadSize,
			ReduceMemoryUsage:     true,
			DisableStartupMessage: true,
			ProxyHeader:           config.App.Server.ProxyHeader,
		})

		// Middlewares
		app.Use(
			cors.New(),
			recover.New(),
			etag.New(),
			compress.New(compress.Config{
				Level: 1,
			}),
		)

		if config.App.Debug {
			app.Use(pprof.New())
		}

		// Register Routes
		router.RegisterRoutes(app)

		// Run application
		app.Listen(config.App.Server.Host + ":" + config.App.Server.Port)
	}
}
