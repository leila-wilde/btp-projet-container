package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// handlers.go - Defines all HTTP request handlers
// these functions receive HTTP requests and return responses

// ## GENRE HANDLERS ##

// GetGenresHandler handles GET /api/genres
// it retrieves all genres and returns them as JSON
func GetGenresHandler(c *fiber.Ctx) error {
	// call our database function to get all genres
	genres, err := GetAllGenres()
	if err != nil {
		// if there's an error, return HTTP 500 (Internal Server Error)
		log.Println("Error in GetGenresHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch genres",
			Code:  500,
		})
	}

	// return genres as JSON with HTTP 200 (OK)
	return c.Status(fiber.StatusOK).JSON(genres)
}

// GetGenreHandler handles GET /api/genres/:id
// it retrieves a single genre by ID
func GetGenreHandler(c *fiber.Ctx) error {
	// c.Params extracts path parameters like :id
	idStr := c.Params("id")

	// strconv.Atoi converts string to integer
	// if it fails, Atoi returns an error
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	// get the genre from database
	genre, err := GetGenreByID(id)
	if err != nil {
		log.Println("Error in GetGenreHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch genre",
			Code:  500,
		})
	}

	// if genre not found, return 404
	if genre == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Genre not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(genre)
}

// CreateGenreHandler handles POST /api/genres
// it creates a new genre from the request body
func CreateGenreHandler(c *fiber.Ctx) error {
	// create a struct to hold the incoming JSON data
	var req struct {
		Name string `json:"name"`
	}

	// c.BodyParser reads the request body and converts JSON into our struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
			Code:  400,
		})
	}

	// validate that name is not empty
	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Name is required",
			Code:  400,
		})
	}

	// create the genre in database
	genre, err := CreateGenre(req.Name)
	if err != nil {
		log.Println("Error in CreateGenreHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to create genre",
			Code:  500,
		})
	}

	// return 201 (Created) with the new genre
	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Message: "Genre created successfully",
		Data:    genre,
	})
}

// UpdateGenreHandler handles PUT /api/genres/:id
// it updates an existing genre
func UpdateGenreHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
			Code:  400,
		})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Name is required",
			Code:  400,
		})
	}

	// update the genre in database
	err = UpdateGenre(id, req.Name)
	if err != nil {
		log.Println("Error in UpdateGenreHandler:", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Genre not found",
			Code:  404,
		})
	}

	// return 200 (OK) with success message
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Genre updated successfully",
	})
}

// DeleteGenreHandler handles DELETE /api/genres/:id
// it deletes a genre
func DeleteGenreHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	err = DeleteGenre(id)
	if err != nil {
		log.Println("Error in DeleteGenreHandler:", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Genre not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Genre deleted successfully",
	})
}

// ## ARTIST HANDLERS ##

func GetArtistsHandler(c *fiber.Ctx) error {
	artists, err := GetAllArtists()
	if err != nil {
		log.Println("Error in GetArtistsHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch artists",
			Code:  500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(artists)
}

func GetArtistHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	artist, err := GetArtistByID(id)
	if err != nil {
		log.Println("Error in GetArtistHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch artist",
			Code:  500,
		})
	}

	if artist == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Artist not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(artist)
}

func CreateArtistHandler(c *fiber.Ctx) error {
	var req struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
			Code:  400,
		})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Name is required",
			Code:  400,
		})
	}

	artist, err := CreateArtist(req.Name)
	if err != nil {
		log.Println("Error in CreateArtistHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to create artist",
			Code:  500,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Message: "Artist created successfully",
		Data:    artist,
	})
}

func UpdateArtistHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
			Code:  400,
		})
	}

	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Name is required",
			Code:  400,
		})
	}

	err = UpdateArtist(id, req.Name)
	if err != nil {
		log.Println("Error in UpdateArtistHandler:", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Artist not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Artist updated successfully",
	})
}

func DeleteArtistHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	err = DeleteArtist(id)
	if err != nil {
		log.Println("Error in DeleteArtistHandler:", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Artist not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Artist deleted successfully",
	})
}

// ## ALBUM HANDLERS ##

func GetAlbumsHandler(c *fiber.Ctx) error {
	albums, err := GetAllAlbums()
	if err != nil {
		log.Println("Error in GetAlbumsHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch albums",
			Code:  500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(albums)
}

func GetAlbumHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	album, err := GetAlbumByID(id)
	if err != nil {
		log.Println("Error in GetAlbumHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch album",
			Code:  500,
		})
	}

	if album == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Album not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(album)
}

func CreateAlbumHandler(c *fiber.Ctx) error {
	var req struct {
		Title    string `json:"title"`
		ArtistID int    `json:"artist_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
			Code:  400,
		})
	}

	if req.Title == "" || req.ArtistID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Title and artist_id are required",
			Code:  400,
		})
	}

	album, err := CreateAlbum(req.Title, req.ArtistID)
	if err != nil {
		log.Println("Error in CreateAlbumHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to create album",
			Code:  500,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Message: "Album created successfully",
		Data:    album,
	})
}

func UpdateAlbumHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	var req struct {
		Title    string `json:"title"`
		ArtistID int    `json:"artist_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
			Code:  400,
		})
	}

	if req.Title == "" || req.ArtistID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Title and artist_id are required",
			Code:  400,
		})
	}

	err = UpdateAlbum(id, req.Title, req.ArtistID)
	if err != nil {
		log.Println("Error in UpdateAlbumHandler:", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Album not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Album updated successfully",
	})
}

func DeleteAlbumHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	err = DeleteAlbum(id)
	if err != nil {
		log.Println("Error in DeleteAlbumHandler:", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Album not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Album deleted successfully",
	})
}

// ## TRACK HANDLERS ##

func GetTracksHandler(c *fiber.Ctx) error {
	tracks, err := GetAllTracks()
	if err != nil {
		log.Println("Error in GetTracksHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch tracks",
			Code:  500,
		})
	}
	return c.Status(fiber.StatusOK).JSON(tracks)
}

func GetTrackHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	track, err := GetTrackByID(id)
	if err != nil {
		log.Println("Error in GetTrackHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to fetch track",
			Code:  500,
		})
	}

	if track == nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Track not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(track)
}

func CreateTrackHandler(c *fiber.Ctx) error {
	var req struct {
		Name    string `json:"name"`
		AlbumID int    `json:"album_id"`
		GenreID int    `json:"genre_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
			Code:  400,
		})
	}

	if req.Name == "" || req.AlbumID == 0 || req.GenreID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Name, album_id, and genre_id are required",
			Code:  400,
		})
	}

	track, err := CreateTrack(req.Name, req.AlbumID, req.GenreID)
	if err != nil {
		log.Println("Error in CreateTrackHandler:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: "Failed to create track",
			Code:  500,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Message: "Track created successfully",
		Data:    track,
	})
}

func UpdateTrackHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	var req struct {
		Name    string `json:"name"`
		AlbumID int    `json:"album_id"`
		GenreID int    `json:"genre_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
			Code:  400,
		})
	}

	if req.Name == "" || req.AlbumID == 0 || req.GenreID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Name, album_id, and genre_id are required",
			Code:  400,
		})
	}

	err = UpdateTrack(id, req.Name, req.AlbumID, req.GenreID)
	if err != nil {
		log.Println("Error in UpdateTrackHandler:", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Track not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Track updated successfully",
	})
}

func DeleteTrackHandler(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
			Code:  400,
		})
	}

	err = DeleteTrack(id)
	if err != nil {
		log.Println("Error in DeleteTrackHandler:", err)
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: "Track not found",
			Code:  404,
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: "Track deleted successfully",
	})
}

// HealthCheckHandler is a simple endpoint to verify the API is running
func HealthCheckHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "OK",
		"message": "InfraMusicStore API is running",
	})
}
