package main

import (
	"database/sql"
	"fmt"
	"log"

	// this imports the MySQL driver - go doesn't automatically include drivers
	// the underscore means we're importing it for its side effects (initialization)
	_ "github.com/go-sql-driver/mysql"
)

// DB is a global variable that holds our database connection
// we make it global so all our functions can use it
var DB *sql.DB

// InitDB initializes the database connection
// this function runs once when the app starts to set up the MySQL connection
func InitDB(dbUser, dbPassword, dbHost, dbPort, dbName string) error {
	// build the connection string (Data Source Name - DSN)
	// format: username:password@tcp(host:port)/database_name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// attempt to open a database connection
	// note: sql.Open doesn't actually connect - it just prepares the connection
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error opening database:", err)
		return err
	}

	// test the connection with a ping
	// this actually tries to connect to verify credentials work
	err = DB.Ping()
	if err != nil {
		log.Println("Error connecting to database:", err)
		return err
	}

	log.Println("✓ Database connected successfully!")
	return nil
}

// ## GENRE FUNCTIONS ##

// GetAllGenres retrieves all genres from the database
func GetAllGenres() ([]Genre, error) {
	// query returns all rows from the genres table
	rows, err := DB.Query("SELECT genre_id, name FROM genres ORDER BY genre_id")
	if err != nil {
		log.Println("Error querying genres:", err)
		return nil, err
	}
	defer rows.Close() // always close rows after we're done

	var genres []Genre // create an empty slice to hold our results

	// loop through each row returned by the query
	for rows.Next() {
		var genre Genre // create a variable to hold this row's data

		// scan copies the row values into our genre variable
		// the order matters - must match the SELECT order
		err := rows.Scan(&genre.GenreID, &genre.Name)
		if err != nil {
			log.Println("Error scanning genre:", err)
			return nil, err
		}

		// append adds the genre to our slice
		genres = append(genres, genre)
	}

	return genres, nil
}

// GetGenreByID retrieves a single genre by its ID
func GetGenreByID(id int) (*Genre, error) {
	var genre Genre

	// QueryRow returns a single row
	// we use LIMIT 1 to ensure we only get one result
	err := DB.QueryRow("SELECT genre_id, name FROM genres WHERE genre_id = ? LIMIT 1", id).
		Scan(&genre.GenreID, &genre.Name)

	// if no rows found, sql.ErrNoRows is returned
	if err == sql.ErrNoRows {
		return nil, nil // return nil to indicate not found
	}
	if err != nil {
		log.Println("Error querying genre:", err)
		return nil, err
	}

	return &genre, nil // return pointer to the genre
}

// CreateGenre creates a new genre in the database
func CreateGenre(name string) (*Genre, error) {
	// Exec runs a query that doesn't return rows (INSERT, UPDATE, DELETE)
	result, err := DB.Exec("INSERT INTO genres (name) VALUES (?)", name)
	if err != nil {
		log.Println("Error creating genre:", err)
		return nil, err
	}

	// LastInsertId gets the ID that was auto-generated for this new row
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return nil, err
	}

	// return the newly created genre
	genre := &Genre{
		GenreID: int(id),
		Name:    name,
	}
	return genre, nil
}

// UpdateGenre updates an existing genre
func UpdateGenre(id int, name string) error {
	// Exec updates the row where genre_id matches
	result, err := DB.Exec("UPDATE genres SET name = ? WHERE genre_id = ?", name, id)
	if err != nil {
		log.Println("Error updating genre:", err)
		return err
	}

	// RowsAffected tells us how many rows were updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	// if 0 rows were affected, the genre doesn't exist
	if rowsAffected == 0 {
		return fmt.Errorf("genre with ID %d not found", id)
	}

	return nil
}

// DeleteGenre deletes a genre from the database
func DeleteGenre(id int) error {
	result, err := DB.Exec("DELETE FROM genres WHERE genre_id = ?", id)
	if err != nil {
		log.Println("Error deleting genre:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("genre with ID %d not found", id)
	}

	return nil
}

// ## ARTIST FUNCTIONS ##

// GetAllArtists retrieves all artists from the database
func GetAllArtists() ([]Artist, error) {
	rows, err := DB.Query("SELECT artist_id, name FROM artists ORDER BY artist_id")
	if err != nil {
		log.Println("Error querying artists:", err)
		return nil, err
	}
	defer rows.Close()

	var artists []Artist

	for rows.Next() {
		var artist Artist
		err := rows.Scan(&artist.ArtistID, &artist.Name)
		if err != nil {
			log.Println("Error scanning artist:", err)
			return nil, err
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

// GetArtistByID retrieves a single artist by ID
func GetArtistByID(id int) (*Artist, error) {
	var artist Artist

	err := DB.QueryRow("SELECT artist_id, name FROM artists WHERE artist_id = ? LIMIT 1", id).
		Scan(&artist.ArtistID, &artist.Name)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("Error querying artist:", err)
		return nil, err
	}

	return &artist, nil
}

// CreateArtist creates a new artist
func CreateArtist(name string) (*Artist, error) {
	result, err := DB.Exec("INSERT INTO artists (name) VALUES (?)", name)
	if err != nil {
		log.Println("Error creating artist:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return nil, err
	}

	artist := &Artist{
		ArtistID: int(id),
		Name:     name,
	}
	return artist, nil
}

// UpdateArtist updates an existing artist
func UpdateArtist(id int, name string) error {
	result, err := DB.Exec("UPDATE artists SET name = ? WHERE artist_id = ?", name, id)
	if err != nil {
		log.Println("Error updating artist:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("artist with ID %d not found", id)
	}

	return nil
}

// DeleteArtist deletes an artist from the database
func DeleteArtist(id int) error {
	result, err := DB.Exec("DELETE FROM artists WHERE artist_id = ?", id)
	if err != nil {
		log.Println("Error deleting artist:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("artist with ID %d not found", id)
	}

	return nil
}

// ## ALBUM FUNCTIONS ##

// GetAllAlbums retrieves all albums from the database
func GetAllAlbums() ([]Album, error) {
	rows, err := DB.Query("SELECT album_id, title, artist_id FROM albums ORDER BY album_id")
	if err != nil {
		log.Println("Error querying albums:", err)
		return nil, err
	}
	defer rows.Close()

	var albums []Album

	for rows.Next() {
		var album Album
		err := rows.Scan(&album.AlbumID, &album.Title, &album.ArtistID)
		if err != nil {
			log.Println("Error scanning album:", err)
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

// GetAlbumByID retrieves a single album by ID
func GetAlbumByID(id int) (*Album, error) {
	var album Album

	err := DB.QueryRow("SELECT album_id, title, artist_id FROM albums WHERE album_id = ? LIMIT 1", id).
		Scan(&album.AlbumID, &album.Title, &album.ArtistID)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("Error querying album:", err)
		return nil, err
	}

	return &album, nil
}

// CreateAlbum creates a new album
func CreateAlbum(title string, artistID int) (*Album, error) {
	result, err := DB.Exec("INSERT INTO albums (title, artist_id) VALUES (?, ?)", title, artistID)
	if err != nil {
		log.Println("Error creating album:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return nil, err
	}

	album := &Album{
		AlbumID:  int(id),
		Title:    title,
		ArtistID: artistID,
	}
	return album, nil
}

// UpdateAlbum updates an existing album
func UpdateAlbum(id int, title string, artistID int) error {
	result, err := DB.Exec("UPDATE albums SET title = ?, artist_id = ? WHERE album_id = ?", title, artistID, id)
	if err != nil {
		log.Println("Error updating album:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("album with ID %d not found", id)
	}

	return nil
}

// DeleteAlbum deletes an album from the database
func DeleteAlbum(id int) error {
	result, err := DB.Exec("DELETE FROM albums WHERE album_id = ?", id)
	if err != nil {
		log.Println("Error deleting album:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("album with ID %d not found", id)
	}

	return nil
}

// ## TRACK FUNCTIONS ##

// GetAllTracks retrieves all tracks from the database
func GetAllTracks() ([]Track, error) {
	rows, err := DB.Query("SELECT track_id, name, album_id, genre_id FROM tracks ORDER BY track_id")
	if err != nil {
		log.Println("Error querying tracks:", err)
		return nil, err
	}
	defer rows.Close()

	var tracks []Track

	for rows.Next() {
		var track Track
		err := rows.Scan(&track.TrackID, &track.Name, &track.AlbumID, &track.GenreID)
		if err != nil {
			log.Println("Error scanning track:", err)
			return nil, err
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

// GetTrackByID retrieves a single track by ID
func GetTrackByID(id int) (*Track, error) {
	var track Track

	err := DB.QueryRow("SELECT track_id, name, album_id, genre_id FROM tracks WHERE track_id = ? LIMIT 1", id).
		Scan(&track.TrackID, &track.Name, &track.AlbumID, &track.GenreID)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("Error querying track:", err)
		return nil, err
	}

	return &track, nil
}

// CreateTrack creates a new track
func CreateTrack(name string, albumID int, genreID int) (*Track, error) {
	result, err := DB.Exec("INSERT INTO tracks (name, album_id, genre_id) VALUES (?, ?, ?)", name, albumID, genreID)
	if err != nil {
		log.Println("Error creating track:", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return nil, err
	}

	track := &Track{
		TrackID: int(id),
		Name:    name,
		AlbumID: albumID,
		GenreID: genreID,
	}
	return track, nil
}

// UpdateTrack updates an existing track
func UpdateTrack(id int, name string, albumID int, genreID int) error {
	result, err := DB.Exec("UPDATE tracks SET name = ?, album_id = ?, genre_id = ? WHERE track_id = ?", name, albumID, genreID, id)
	if err != nil {
		log.Println("Error updating track:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("track with ID %d not found", id)
	}

	return nil
}

// DeleteTrack deletes a track from the database
func DeleteTrack(id int) error {
	result, err := DB.Exec("DELETE FROM tracks WHERE track_id = ?", id)
	if err != nil {
		log.Println("Error deleting track:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("track with ID %d not found", id)
	}

	return nil
}
