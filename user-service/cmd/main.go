package main

import (
	"database/sql"
	"fmt"
	"log"
	"use/internal/api"
	"use/internal/repository"
	"use/internal/service"

	use "use/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"use/config"
	migration "use/db"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// Connect to the database
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migration.RunMigration(db)
	// Initialize the repository and service layers
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	// Initialize the API
	userAPI := api.NewUserAPI(userService)

	// Initialize the HTTP server
	httpAddr := ":" + conf.HttpPort
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// Define HTTP routes
	server.GET("/users", userAPI.ListUsers)
	server.GET("/users/:id", userAPI.GetUser)
	server.POST("/users", userAPI.CreateUser, use.UserValidator)
	server.PUT("/users/:id", userAPI.UpdateUser, use.UserValidator)
	server.DELETE("/users/:id", userAPI.DeleteUser)

	// Start the HTTP server
	log.Printf("Starting HTTP server on %s", httpAddr)
	err = server.Start(httpAddr)
	if err != nil {
		log.Fatal(err)
	}
}
