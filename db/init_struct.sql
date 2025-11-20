/* Drop database if it already exists */

DROP DATABASE IF EXISTS inframusic_db;

/* Create database */

CREATE DATABASE inframusic_db;

USE inframusic_db;

/* Create tables for the inframusic_db database */

CREATE TABLE artists (
    artist_id INT NOT NULL,
    name VARCHAR(255),
    CONSTRAINT PK_artist PRIMARY KEY (artist_id)
);

CREATE TABLE albums (
    album_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    artist_id INT NOT NULL,
    CONSTRAINT PK_album PRIMARY KEY (album_id)
);

CREATE TABLE tracks (
    track_id INT NOT NULL,
    album_id INT,
    genre_id INT,
    name VARCHAR(255) NOT NULL,
    CONSTRAINT PK_tracks PRIMARY KEY (track_id)
);

CREATE TABLE genres (
    genre_id INT NOT NULL,
    name VARCHAR(255),
    CONSTRAINT PK_genre PRIMARY KEY (genre_id)
);