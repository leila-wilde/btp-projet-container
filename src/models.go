package main

// models.go - defines the data structures for our API
// these "structs" are similar to classes in Java or Python - they define the shape of our data

// Genre represents a music genre
// `json:"field_name"` tells Go how to convert to/from JSON when we send responses
type Genre struct {
	GenreID int    `json:"genre_id"` // ID is automatically converted from int to JSON number
	Name    string `json:"name"`
}

// Artist represents a music artist/band
type Artist struct {
	ArtistID int    `json:"artist_id"`
	Name     string `json:"name"`
}

// Album represents a music album
type Album struct {
	AlbumID  int    `json:"album_id"`
	Title    string `json:"title"`
	ArtistID int    `json:"artist_id"`
}

// Track represents a music track (song)
type Track struct {
	TrackID int    `json:"track_id"`
	Name    string `json:"name"`
	AlbumID int    `json:"album_id"`
	GenreID int    `json:"genre_id"`
}

// ErrorResponse is used when we need to return an error to the client
// this gives consistent error messages across our API
type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// SuccessResponse is used for create/update operations
// it confirms the operation was successful
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // interface{} means any type can go here
}
