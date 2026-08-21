package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UrlRequestBody struct {
	Url string `json:"url"`
}

func createUrl(w http.ResponseWriter, r *http.Request) {

	var requestBody UrlRequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, repoErr := CreateUrl(requestBody.Url)
	if repoErr != nil {
		if repoErr.Code == ErrCodeAlreadyExists {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Println(repoErr.Message)
		return
	}

	j, _ := json.Marshal(url)
	fmt.Fprint(w, string(j))
}

func updateUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("url")

	_, err := GetUrlByShort(shortUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

	}

	var requestBody UrlRequestBody
	decodeError := json.NewDecoder(r.Body).Decode(&requestBody)
	if decodeError != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedUrl, err := UpdateUrl(shortUrl, requestBody.Url)

	j, _ := json.Marshal(updatedUrl)
	fmt.Fprint(w, string(j))
}

func deleteUrl(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("url")

	err := DeleteUrl(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("url")

	url, err := GetUrlByShort(shortUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	j, _ := json.Marshal(url)
	fmt.Fprint(w, string(j))
}

func fetchAllUrls(w http.ResponseWriter, r *http.Request) {

	data, _ := json.Marshal(shortenUrls)

	fmt.Fprint(w, string(data))
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /shorten/{url}", loggingMiddleware(getUrl))
	mux.HandleFunc("POST /shorten", loggingMiddleware(createUrl))
	mux.HandleFunc("PUT /shorten/{url}", loggingMiddleware(updateUrl))
	mux.HandleFunc("DELETE /shorten/{url}", loggingMiddleware(deleteUrl))
	mux.HandleFunc("/all", loggingMiddleware(fetchAllUrls))

	s := &http.Server{
		Addr:    ":8585",
		Handler: mux,
	}

	log.Printf("Server is up and listening on http://localhost%s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
