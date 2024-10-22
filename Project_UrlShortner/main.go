package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID string `json:"id"`
	OriginalURL string `json:"original_url"`
	ShortURL string `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) // convert string to slice of bytes
	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	return hash[:8] // return only the first 8 characters
}

func createUrl(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL
	urlDB[id] = URL{
		ID: id,
		OriginalURL: OriginalURL,
		ShortURL: shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL ,string){
	url ,ok := urlDB[id]
	if !ok {
		return URL{}, "url not found"
	}
	return url, ""
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct{
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	shortURL_ := createUrl(data.URL)
	// fmt.Printf("Short URL: %s\n", shortURL)
	// fmt.Fprintf(w, shortURL)
	response := struct{
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != "" {
		http.Error(w, err, http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	// fmt.Println("Making a url shortner app using golang")
	// OriginalURL := "https://www.google.com"
	// shortURL := generateShortURL(OriginalURL)

	// register the handler function to handle all requests to the root URL("/")
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)
	 
	// start the HTTP server on port 8000
	fmt.Println("Starting the HTTP server on port 8000")
	err :=http.ListenAndServe(":8000", nil) 
	if err != nil {
		fmt.Println("Error starting the HTTP server" , err)
	}
}