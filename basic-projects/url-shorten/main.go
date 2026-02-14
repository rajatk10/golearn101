package main

import (
	"log"
	"net/http"
)

func main() {
	if err := loadURLStoreFile(); err != nil {
		log.Printf("Failed to load URL store File: %v", err)
	}
	startAutoSave()
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/sh/", redirectHandler)
	log.Println("Server started at :30805")
	if err := http.ListenAndServe(":30805", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
