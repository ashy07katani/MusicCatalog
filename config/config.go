package config

import (
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type LocalConfig struct {
	SpotifyKey      string        `env:"SPOTIFY_KEY"`
	SpotifySecret   string        `env:"SPOTIFY_SECRET"`
	SpotifyURL      string        `env:"SPOTIFY_URL"`
	RedisHost       string        `env:"REDIS_HOST"`
	RedisPort       int           `env:"REDIS_PORT"`
	RedisDefaultDB  int           `env:"REDIS_DEFUALTDB"`
	CacheExpiryTime time.Duration `env:"CACHEEXPIRYTIME"`
}

func GetConfig() (*LocalConfig, error) {
	var localConfig LocalConfig
	//will use godotenv to load the file
	err := godotenv.Load("./local.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	//getting the expiry time for cache and converting it to time.Duration instance
	cacheETStr := os.Getenv("CACHEEXPIRYTIME")
	cacheET, err := time.ParseDuration(cacheETStr)
	if err != nil {
		log.Fatalf("Invalid CACHE_ET %v", err)
	}
	//from the loaded file we will parse the environment variable into localconfig instance
	if err := env.Parse(&localConfig); err != nil {
		log.Fatalf("Error reading the environment variable %v", err)
		return nil, err
	}
	//setting the cache Expiry time
	localConfig.CacheExpiryTime = cacheET
	log.Printf("%+v\n", localConfig)
	return &localConfig, nil
}
