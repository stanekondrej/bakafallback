package main

import (
	"log"
	"net/url"
	"os"

	"github.com/stanekondrej/bakafallback/internal/app/bakafallback"
	"github.com/stanekondrej/bakalari/pkg/bakalari"
)

func getEnv() (username, password, url string) {
	username, ok := os.LookupEnv("BAKAFALLBACK_USERNAME")
	if !ok {
		log.Fatal("BAKAFALLBACK_USERNAME není nastaveno")
	}

	password, ok = os.LookupEnv("BAKAFALLBACK_PASSWORD")
	if !ok {
		log.Fatal("BAKAFALLBACK_PASSWORD není nastaveno")
	}

	url, ok = os.LookupEnv("BAKAFALLBACK_URL")
	if !ok {
		log.Fatal("BAKAFALLBACK_URL není nastaveno")
	}

	return
}

func main() {
	username, password, u := getEnv()
	parsedUrl, err := url.Parse(u)
	if err != nil {
		log.Fatal("Adresa je neplatná")
	}

	api := bakalari.NewApi(parsedUrl)
	accessToken, refreshToken, err := api.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Přihlášen jako uživatel %s", username)

	bakafallback.StartServer(accessToken, refreshToken, api)
}
