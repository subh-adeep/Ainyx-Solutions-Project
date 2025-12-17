package main

import (
	"database/sql"
	"errors"
	"project/config"
	db "project/db/sqlc"
	"project/internal/handler"
	"project/internal/logger"
	"project/internal/repository"
	"project/internal/routes"
	"project/internal/service"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	// Initialize Config
	cfg := config.LoadConfig()

	// Initialize Logger
	logger.InitLogger()
	defer logger.Log.Sync()

	logger.Log.Info("Starting application...")

	// Connect to Database
	conn, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		logger.Log.Fatal("cannot connect to db", zap.Error(err))
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		logger.Log.Fatal("cannot ping db", zap.Error(err))
	}

	logger.Log.Info("Connected to database")

	// Initialize Layers
	queries := db.New(conn)
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Initialize Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			logger.Log.Error("Fiber Error", zap.Error(err))
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Setup Routes
	routes.SetupRoutes(app, userHandler)

	// Start Server
	port := cfg.Port
	logger.Log.Info("Server listening", zap.String("port", port))
	if err := app.Listen(":" + port); err != nil {
		logger.Log.Fatal("Server failed", zap.Error(err))
	}
}
