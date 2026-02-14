package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var urlShortFile = "urlstore.json"

type URLData struct {
	LongURL  string `json:"url"`
	ShortURL string `json:"short_url"`
}

var (
	urlstore = make(map[string]string)
	mu       sync.Mutex
)

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req URLData
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if _, err := url.ParseRequestURI(req.LongURL); err != nil {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}
	short := generateShortURL()
	mu.Lock()
	for {
		if _, exists := urlstore[short]; !exists {
			break
		}
		short = generateShortURL()
	}
	urlstore[short] = req.LongURL
	mu.Unlock()
	//Save the shortened data to file
	if err := SaveURLShortStoreFile(); err != nil {
		log.Printf("Failed to save URL store File: %v", err)
	}
	resp := URLData{LongURL: req.LongURL, ShortURL: fmt.Sprintf("http://localhost:30805/sh/%s", short)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	short := r.URL.Path[len("/sh/"):]
	mu.Lock()
	longURL, exists := urlstore[short]
	mu.Unlock()
	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}
func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Fatalf("Failed to generate random index: %v", err)
		}
		b[i] = charset[randomIndex.Int64()]
	}
	log.Println("Shorten URL code: ", string(b))
	return string(b)
}

func loadURLStoreFile() error {
	data, err := os.ReadFile(urlShortFile)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(urlShortFile)
			return nil
		}
		log.Fatalf("Failed to read URL store: %v", err)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(data) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, &urlstore); err != nil {
		log.Fatalf("Failed to parse URL store: %v", err)
	}
	return nil
}

func SaveURLShortStoreFile() error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.MarshalIndent(urlstore, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(urlShortFile, data, 0644)
}

func startAutoSave() {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			if err := SaveURLShortStoreFile(); err != nil {
				log.Printf("Failed to save URL store File: %v", err)
			}
		}
	}()
}
