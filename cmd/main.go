package main

import (

	// Load .env file
	_ "express_be/docs"
	"express_be/server"
	"flag"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// @title CMN Express API Documentation
// @version 1.0
// @description Welcome to the CMN Express API documentation. This server facilitates seamless integration for managing logistics, tracking, and delivery services efficiently.
// @description Explore the endpoints to unlock robust features designed to enhance your operational workflow.
// @contact.name Lucas
// @contact.url https://sharkytech.vercel.app
// @contact.email nguyenmanh180102@gmail.com
// @contact.description For technical inquiries or support, feel free to reach out to our dedicated team.

// @host 203.145.47.225
// @BasePath /
// @termsOfService https://sharkytech.vercel.app/terms
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

func main() {
	// Dùng cách cờ để đọc file config
	// go run ./cmd/main.go -config="./conf/pgsql.env"
	configPathFlag := flag.String("config", "", "Path to configuration file")
	flag.Parse()
	// export CONFIG_PATH="./conf/pgsql.env"
	// go run main.go
	configPath := *configPathFlag

	log.Printf("Starting server with config: %s", configPath)
	server.Run(configPath)
}
