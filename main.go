package main

import (
	"log"
	"musiccatalog/config"
	"musiccatalog/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	localConfig, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error initializing  configurations %v", err)
		return
	}
	log.Printf("%+v", localConfig)
	catalogService := service.NewService(localConfig.SpotifyKey, localConfig.SpotifySecret, localConfig.SpotifyURL)
	router := mux.NewRouter()
	router.HandleFunc("/api/spotify/accessToken", catalogService.GetSpotifyAccessToken)
	http.ListenAndServe(":9000", router)
}
