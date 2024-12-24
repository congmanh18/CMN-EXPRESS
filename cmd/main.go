package main

import (

	// Load .env file
	"express_be/server"
	"flag"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Dùng cách cờ để đọc file config
	// go run ./cmd/main.go -config="./conf/pgsql.env"
	configPathFlag := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	// export CONFIG_PATH="./conf/pgsql.env"
	// go run main.go
	configPath := *configPathFlag
	// if configPath == "" {
	// 	configPath = os.Getenv("CONFIG_PATH")
	// }
	// if configPath == "" {
	// 	configPath = "D://Go-Basic//nextedu//docs-service//deploy//local_conf.env"
	// }

	log.Printf("Starting server with config: %s", configPath)
	server.Run(configPath)
}

// Cách 1:
// go run main.go -config="./docs-service/deploy/prod_conf.env"

// Cách 2:
// export CONFIG_PATH="./docs-service/deploy/prod_conf.env"
// go run main.go
