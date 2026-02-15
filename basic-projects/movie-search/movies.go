package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Config struct {
	OmdbApiKey string `json:"omdb_api_key"`
	DbName     string `json:"db_name"`
}

type Movie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Director   string `json:"Director"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
	Response   string `json:"Response"`
	Error      string `json:"Error,omitempty"`
}

func loadConfig() (*Config, error) {
	//File → Read all bytes → Unmarshal → Struct
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Println("Unable to json keys from config.json")
		return nil, err
	}
	return &config, nil
}

func searchMovie(title string) (*Movie, error) {
	baseUrl := "http://www.omdbapi.com/"
	params := url.Values{}
	config, err := loadConfig()
	if err != nil {
		log.Println("Unable to read config file, api key")
		return nil, err
	}
	apiKey := config.OmdbApiKey
	params.Add("apikey", apiKey)
	params.Add("t", title)
	fullUrl := baseUrl + "?" + params.Encode()
	log.Println("Here is full URL : ", fullUrl)

	//Make http call to omdb
	resp, err := http.Get(fullUrl)
	// resp.Body is an io.Reader (stream, not []byte)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//HTTP response → Stream body → Decode → Struct

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}

	if movie.Response == "False" {
		return nil, fmt.Errorf("movie not found")
	}
	return &movie, nil

}
