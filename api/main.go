package main

import (
	"crypto/subtle"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"website-api/internal/database"
	"website-api/internal/handlers"
	"website-api/pkg"
)

type jwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	// Load .env file
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect to database
	db, err := database.NewDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close() // Close database connection when main function ends

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// public routes
	e.Static("/", "web")
	e.File("/admin", "web/admin.html")

	api := e.Group("/api")

	// Health check
	api.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Erthang's working just fine!")
	})

	// Basic auth middleware
	api.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(cfg.APIUsername)) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte(cfg.APIPassword)) == 1 {
			return true, nil
		}
		return false, nil
	}))

	projectHandler := handlers.NewProjectHandler(db)
	api.GET("/projects", projectHandler.GetProjects)
	api.GET("/projects/list", projectHandler.GetProjectList)
	api.GET("/projects/:name", projectHandler.GetProject)
	api.PUT("/projects", projectHandler.EditProject)
	api.DELETE("/projects/:name", projectHandler.DeleteProject)
	api.POST("/projects", projectHandler.CreateProject)

	// Start the server
	address := "0.0.0.0:" + cfg.PORT
	e.Logger.Fatal(e.Start(address))
}
