package main

import (
	"fmt"
	"log"
	"os"
	"project-url-shortner/db"
	"project-url-shortner/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
type Config struct{
	Host string
	Port string
	User string		
	Password string
	SSLMode string
	DBName string
}	
func main() {
err := godotenv.Load(".env"); if err != nil{
		log.Fatal((err))
	}
	config := &Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password:os.Getenv("DB_PASSWORD"),		
		SSLMode:os.Getenv("DB_SSLMODE"),
		DBName:os.Getenv("DB_DBNAME"),
	}
	fmt.Println(config)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	fmt.Println(connectionString)
	// Initialize the database
	err = db.InitDB(connectionString); if err != nil {
		fmt.Println("Could not connect to the database", err)
		return
	}
	// Initialize the Gin router
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}