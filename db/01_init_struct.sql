/* Drop database if it already exists */

DROP DATABASE IF EXISTS inframusic_db;

/* Create database */

CREATE DATABASE inframusic_db;

USE inframusic_db;

/* Create tables for the inframusic_db database */

CREATE TABLE genres (
    genre_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    CONSTRAINT PK_genre PRIMARY KEY (genre_id)
);

CREATE TABLE artists (
    artist_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    CONSTRAINT PK_artist PRIMARY KEY (artist_id)
);

CREATE TABLE albums (
    album_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist_id INT NOT NULL,
    -- CONSTRAINT PK_album PRIMARY KEY (album_id),
    CONSTRAINT FK_artist_id FOREIGN KEY (artist_id) REFERENCES artists (artist_id)
);

CREATE TABLE tracks (
    track_id INT NOT NULL AUTO_INCREMENT,
    album_id INT,
    genre_id INT,
    name VARCHAR(255) NOT NULL,
    CONSTRAINT PK_tracks PRIMARY KEY (track_id),
    CONSTRAINT FK_genres FOREIGN KEY (genre_id) REFERENCES genres (genre_id),
    CONSTRAINT FK_albums FOREIGN KEY (album_id) REFERENCES albums (album_id)
);

