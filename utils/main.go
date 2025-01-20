package main

import (
	"github.com/go-chi/cors"

	handler "express_be/handler/message"
	"express_be/provider"
	"express_be/repository/message_service/conversation"
	"express_be/repository/message_service/message"
	"express_be/repository/message_service/participant"
	"express_be/server"

	usecase "express_be/usecase/chat"
	"express_be/utils/chat/config"
	"express_be/utils/chat/logger"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	socketio "github.com/nkovacs/go-socket.io"
)

func main() {
	// Load configuration
	conf := config.GetInstance()
	if !conf.Load("utils/configs/config.json") {
		log.Fatal("Failed to load configuration")
	}

	// Load service environment variables
	serviceConf := server.LoadServiceConfig("local/local.env")
	appProvider := provider.NewAppProvider(serviceConf)

	// Setup logger
	logger.SetupLogger(conf.Logging.Enable, conf.Logging.Level)

	// Initialize repositories
	conversationRepo := conversation.NewRepo(appProvider.Postgres)
	messageRepo := message.NewRepo(appProvider.Postgres)
	participantRepo := participant.NewRepo(appProvider.Postgres)

	// Initialize use cases
	chatUseCase := usecase.NewChatUseCase(conversationRepo, messageRepo, participantRepo)

	// Initialize SocketHandler
	socketHandler := handler.NewSocketHandler(chatUseCase)

	// Create Socket.IO server
	socketServer, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatalf("Failed to create Socket.IO server: %v", err)
	}

	// Setup routes for Socket.IO
	socketHandler.SetupRoutes(socketServer)

	// Initialize Chi router
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Hoặc thay "*" bằng nguồn gốc cụ thể, ví dụ: "http://localhost:3000"
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	r.Use(cors.Handler)
	r.Use(middleware.Logger)    // Request logging
	r.Use(middleware.Recoverer) // Recover from panics
	// Serve static files (optional)
	r.Handle("/", http.FileServer(http.Dir("./assets")))

	// Attach Socket.IO server to HTTP server
	r.Mount("/socket.io", http.StripPrefix("/socket.io", socketServer))

	// Start HTTP server
	addr := fmt.Sprintf("%s:%d", conf.ChatConfig.Host, conf.ChatConfig.Port)
	log.Printf("Server is running at %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
