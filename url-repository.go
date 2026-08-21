package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type ShortURL struct {
	Id        string `json:"id"`
	LongUrl   string `json:"longUrl"`
	ShortUrl  string `json:"shortUrl"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

const (
	ErrCodeNotFound      = 1
	ErrCodeAlreadyExists = 2
)

type RepositoryError struct {
	Code    int
	Message string
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("error code %d: %s", e.Code, e.Message)
}

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

var shortenUrls []ShortURL

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetUrlByLong(longUrl string) (ShortURL, *RepositoryError) {
	for _, url := range shortenUrls {
		if url.LongUrl == longUrl {
			return url, nil
		}
	}

	return ShortURL{}, &RepositoryError{
		Code:    ErrCodeNotFound,
		Message: "URL not found",
	}
}

func GetUrlByShort(shortUrl string) (ShortURL, *RepositoryError) {
	for _, url := range shortenUrls {
		if url.ShortUrl == shortUrl {
			return url, nil
		}
	}

	return ShortURL{}, &RepositoryError{
		Code:    ErrCodeNotFound,
		Message: "URL not found",
	}
}

func CreateUrl(longUrl string) (ShortURL, *RepositoryError) {
	url, err := GetUrlByLong(longUrl)
	if err == nil {
		return url, &RepositoryError{
			Code:    ErrCodeAlreadyExists,
			Message: "URL already exist",
		}
	}

	newUrl := ShortURL{
		Id:        fmt.Sprint(len(shortenUrls)),
		LongUrl:   longUrl,
		ShortUrl:  generateRandomString(8),
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	shortenUrls = append(shortenUrls, newUrl)

	return newUrl, nil
}

func UpdateUrl(shortUrl, newLongUrl string) (ShortURL, *RepositoryError) {
	for i, url := range shortenUrls {
		if url.ShortUrl == shortUrl {
			log.Println("Found URL", url)
			shortenUrls[i].LongUrl = newLongUrl
			shortenUrls[i].UpdatedAt = time.Now().UTC().Format(time.RFC3339)
			return shortenUrls[i], nil
		}
	}

	return ShortURL{}, &RepositoryError{
		Code:    ErrCodeNotFound,
		Message: "URL doesn't exist",
	}
}

func DeleteUrl(shortUrl string) *RepositoryError {
	for i, url := range shortenUrls {
		if url.ShortUrl == shortUrl {
			shortenUrls = append(shortenUrls[:i], shortenUrls[i+1:]...)
			return nil
		}
	}

	return &RepositoryError{
		Code:    ErrCodeNotFound,
		Message: "URL doesn't exist",
	}
}
