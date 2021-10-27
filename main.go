package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"team1.asia/fibo/config"
	"team1.asia/fibo/db"
	"team1.asia/fibo/log"
	"team1.asia/fibo/router"
)

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
		app.Use(recover.New())
		app.Use(etag.New())
		app.Use(compress.New(compress.Config{
			Level: 1,
		}))

		app.Use(logger.New(logger.Config{
			Format:     "${pid} ${status} - ${method} ${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "America/New_York",
		}))

		if config.App.Debug {
			app.Use(pprof.New())
		}

		// Register Routes
		router.RegisterRoutes(app)

		// Run application
		app.Listen(config.App.Server.Host + ":" + config.App.Server.Port)
	}
}
