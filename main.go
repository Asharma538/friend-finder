package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/Asharma538/friend-finder/handlers"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	// Determine if running in CLI mode or server mode
	if len(os.Args) > 1 && os.Args[1] == "server" {
		startWebServer()
	} else {
		startCLI()
	}
}

func startWebServer() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/api/check-username", handlers.CheckUsernameHandler)

	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	fmt.Printf("🚀 Friend Finder server starting on http://localhost:%s\n", port)
	fmt.Printf("📁 Serving static files from ./static/\n")
	fmt.Printf("🔍 API endpoint: http://localhost:%s/api/check-username\n", port)
	fmt.Println("Press Ctrl+C to stop the server")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func startCLI() {
	fmt.Print("Welcome to the Free Username Finder! \n You can check the availability of usernames across multiple platforms at once.\n\n")
	fmt.Println("💡 Tip: Run './friend-finder.exe server' to start the web interface!")
	fmt.Println()

	// Import the CLI functionality
	runCLI()
}
