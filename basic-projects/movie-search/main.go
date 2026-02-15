package main

import (
	"flag"
	"log"
)

const movieDBFile string = "data/movies.db"

func main() {
	title := flag.String("title", "", "Movie title")
	flag.Parse()

	if *title == "" {
		log.Println("Please provide a movie title to search")
		return
	}
	//Initialize DB
	db, err := initDB(movieDBFile)
	if err != nil {
		log.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	//Search Movie
	movie, err := searchMovieWithCache(db, *title)
	if err != nil {
		log.Println("Error searching movie:", err)
		return
	}
	log.Printf("Title : %v\n", movie.Title)
	log.Printf("Year : %v\n", movie.Year)
	log.Printf("Rating : %v\n", movie.ImdbRating)
	log.Printf("Director : %v\n", movie.Director)
	log.Printf("Actors : %v\n", movie.Actors)
	log.Printf("Plot : %v\n", movie.Plot)
}
