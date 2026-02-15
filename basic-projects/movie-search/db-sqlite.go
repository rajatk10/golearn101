package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const MovieTable string = `
	CREATE TABLE IF NOT EXISTS movies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    imdb_id TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    year TEXT,
    director TEXT,
    actors TEXT,
    plot TEXT,
    imdb_rating TEXT,
    response TEXT,
    fetched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

func initDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	//Test DB connection
	if err := db.Ping(); err != nil {
		log.Println("Unable to connect to database")
		return nil, err
	}

	//Create table if not exists
	if _, err := db.Exec(MovieTable); err != nil {
		log.Println("Unable to create movie DB Table")
		return nil, err
	}
	return db, nil
}

func insertMovie(db *sql.DB, movie *Movie) error {
	//SQL Query
	query := `INSERT INTO movies (imdb_id, title, year, director, actors, plot, imdb_rating, response) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(
		query,
		movie.ImdbID,
		movie.Title,
		movie.Year,
		movie.Director,
		movie.Actors,
		movie.Plot,
		movie.ImdbRating,
		movie.Response,
	)
	if err != nil {
		return err
	}
	return nil
}

func getMovieFromDB(db *sql.DB, title string) (*Movie, error) {
	//SQL Query
	query := `
			  SELECT imdb_id, title, year, director, actors, plot, imdb_rating, response 
			  FROM movies 
			  WHERE LOWER(title) = LOWER(?)
    `

	var movie Movie

	err := db.QueryRow(query, title).Scan(
		&movie.ImdbID,
		&movie.Title,
		&movie.Year,
		&movie.Director,
		&movie.Actors,
		&movie.Plot,
		&movie.ImdbRating,
		&movie.Response,
	)

	if err == sql.ErrNoRows {
		log.Printf("No movie with title %s", title)
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &movie, nil

}

func searchMovieWithCache(db *sql.DB, title string) (*Movie, error) {
	//Check cache
	cached, err := getMovieFromDB(db, title)
	if cached != nil {
		log.Println("Found in cache ! ", strings.ToUpper(title))
		return cached, nil
	}

	log.Println("Not in cache, fetching from API.... ", strings.ToUpper(title))
	movie, err := searchMovie(title)
	if err != nil {
		return nil, err
	}
	if err := insertMovie(db, movie); err != nil {
		log.Println("Warning: Unable to insert movie into DB - cache failed : ", strings.ToUpper(title))
	} else {
		log.Println("Movie cached successfully ! ", strings.ToUpper(title))
	}
	return movie, nil
}
