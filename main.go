package main

import (
	"log"
	"musiccatalog/config"
	"musiccatalog/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func main() {
	localConfig, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error initializing  configurations %v", err)
		return
	}
	log.Printf("%+v", localConfig)
	catalogService := service.NewService(localConfig.SpotifyKey, localConfig.SpotifySecret, localConfig.SpotifyURL)
	redisService := service.NewRedisService(localConfig.RedisPort, localConfig.RedisHost, localConfig.RedisDefaultDB, localConfig.CacheExpiryTime)
	rDB := redis.NewClient(&redis.Options{
		Addr:     redisService.Host + ":" + strconv.Itoa(redisService.RedisPort),
		Password: "",
		DB:       redisService.DefaultDB,
	})
	musicCatalogService := service.MusicCatalogService{
		RedisClient:     rDB,
		CacheExpiryTime: redisService.CacheExpiryTime,
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/spotify/accessToken", catalogService.GetSpotifyAccessToken)
	router.HandleFunc("/api/songs", musicCatalogService.GetSong)
	http.ListenAndServe(":9000", router)
}
