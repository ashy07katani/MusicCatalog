package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"strings"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	RedisPort       int
	Host            string
	DefaultDB       int
	CacheExpiryTime time.Duration
}
type MusicCatalogService struct {
	RedisClient     *redis.Client
	CacheExpiryTime time.Duration
}

func NewRedisService(redisPort int, host string, defaultDB int, cacheET time.Duration) *RedisService {
	return &RedisService{
		RedisPort:       redisPort,
		Host:            host,
		DefaultDB:       defaultDB,
		CacheExpiryTime: cacheET,
	}
}

func (c *MusicCatalogService) getSongFromCache(songName string, ctx context.Context) (string, bool) {
	log.Println("inside getSongFromCache")
	val, err := c.RedisClient.Get(ctx, songName).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		log.Fatalf("Error fetching from Cache: %v", err)
		return "", false
	}
	return val, true
}

func (c *MusicCatalogService) setSongToCache(ctx context.Context, songName string, value string, cacheET time.Duration) bool {
	err := c.RedisClient.Set(ctx, songName, value, cacheET)
	if err != nil {
		return false
	}
	return true
}

func (c *MusicCatalogService) fetchSongFromAPI(songName string) (map[string]interface{}, error) {
	songDetail := map[string]interface{}{
		"id":          67890,
		"title":       "Sacrifice",
		"artist":      "The Weeknd",
		"album":       "Dawn FM",
		"duration":    195, // duration in seconds
		"genre":       "R&B",
		"releaseYear": 2022,
		"rating":      4.7, // out of 5
		"isFavorite":  true,
		"lyrics":      "I was born in a city where the winter nights don't ever sleep...",
	}
	log.Println("Inside fetchSongFromAPI", songName)
	return songDetail, nil
}

func (c *MusicCatalogService) GetSong(rw http.ResponseWriter, r *http.Request) {
	//check if the song is present in the cache.
	query := r.URL.Query()
	songName := strings.ToLower(query.Get("songname"))
	songNameFormatted := strings.ReplaceAll(songName, " ", "")
	ctx := context.Background()
	song, found := c.getSongFromCache(songNameFormatted, ctx)

	// if present in the cache then retieve the song
	if found {
		log.Println("The fetched song from cache is ", song)
		rw.Write([]byte(song))
		return
	}

	//else make a request to spotify
	songData, err := c.fetchSongFromAPI(songName)
	if err != nil {
		log.Println("Failed to fetch data from API", err)
	}
	log.Println("The fetched song from API call is ", songData)
	//convert the data to json
	jsonData, err := json.Marshal(songData)
	if err != nil {
		log.Fatalf("Cannot convert the given data to Json %v", jsonData)
	}

	//set the response to the cache
	c.setSongToCache(ctx, songNameFormatted, string(jsonData), c.CacheExpiryTime)
	rw.Write(jsonData)
}
