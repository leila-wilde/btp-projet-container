package main

import (
	"log"
	"os"

	// Fiber is a web framework for Go - similar to Express in Node.js
	"github.com/gofiber/fiber/v2"
	// this middleware adds CORS headers to allow cross-origin requests
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// main.go - entry point of our application
// the main() function runs when we start the program

func main() {
	// ## DATABASE SETUP ##

	// get database configuration from environment variables
	// environment variables are set in docker-compose.yml or .env file
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	apiPort := os.Getenv("API_PORT")

	// set defaults if environment variables aren't set
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbUser == "" {
		dbUser = "inframusic_user"
	}
	if dbPassword == "" {
		dbPassword = "inframusic_pass"
	}
	if dbName == "" {
		dbName = "inframusic_db"
	}
	if apiPort == "" {
		apiPort = "8080"
	}

	// initialize the database connection
	// this opens a connection to MySQL
	err := InitDB(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
		// log.Fatal stops the program immediately if database fails
	}

	// ## FIBER APP SETUP ##

	// create a new Fiber application
	// Fiber is an Express-like HTTP framework - it handles routing and HTTP stuff
	app := fiber.New(fiber.Config{
		Prefork: false, // set to false for development (set true in production for better performance)
	})

	// add CORS middleware - allows requests from different domains
	// this is important for frontend apps making requests to our API
	app.Use(cors.New())

	// ## ROUTES ##

	// health check endpoint - simple way to verify API is running
	app.Get("/api/health", HealthCheckHandler)

	// ## GENRE ROUTES ##
	// GET /api/genres - list all genres
	app.Get("/api/genres", GetGenresHandler)

	// GET /api/genres/:id - get a single genre by ID
	// :id is a path parameter that gets extracted and passed to the handler
	app.Get("/api/genres/:id", GetGenreHandler)

	// POST /api/genres - create a new genre
	app.Post("/api/genres", CreateGenreHandler)

	// PUT /api/genres/:id - update an existing genre
	app.Put("/api/genres/:id", UpdateGenreHandler)

	// DELETE /api/genres/:id - delete a genre
	app.Delete("/api/genres/:id", DeleteGenreHandler)

	// ## ARTIST ROUTES ##
	app.Get("/api/artists", GetArtistsHandler)
	app.Get("/api/artists/:id", GetArtistHandler)
	app.Post("/api/artists", CreateArtistHandler)
	app.Put("/api/artists/:id", UpdateArtistHandler)
	app.Delete("/api/artists/:id", DeleteArtistHandler)

	// ## ALBUM ROUTES ##
	app.Get("/api/albums", GetAlbumsHandler)
	app.Get("/api/albums/:id", GetAlbumHandler)
	app.Post("/api/albums", CreateAlbumHandler)
	app.Put("/api/albums/:id", UpdateAlbumHandler)
	app.Delete("/api/albums/:id", DeleteAlbumHandler)

	// ## TRACK ROUTES ##
	app.Get("/api/tracks", GetTracksHandler)
	app.Get("/api/tracks/:id", GetTrackHandler)
	app.Post("/api/tracks", CreateTrackHandler)
	app.Put("/api/tracks/:id", UpdateTrackHandler)
	app.Delete("/api/tracks/:id", DeleteTrackHandler)

	// ## START SERVER ##

	// log that the server is starting
	log.Printf("InfraMusicStore API starting on port %s...", apiPort)

	// start listening for HTTP requests on the specified port
	// this# GEN blocks until the server is stopped
	err = app.Listen(":" + apiPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
