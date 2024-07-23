package service

import (
	b64 "encoding/base64"
	"encoding/json"
	"log"
	"musiccatalog/dto"
	"musiccatalog/utilities"
	"net/http"
	"net/url"
)

type Service struct {
	clientId string
	secret   string
	url      string
}

func NewService(clientId string, secret string, url string) *Service {
	return &Service{clientId: clientId,
		secret: secret,
		url:    url}
}
func (service *Service) GetSpotifyAccessToken(rw http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	parsedURL, err := url.Parse(service.url)
	if err != nil {
		log.Fatalf("Error in parsing request URL to hit Spotify endpoint %v", err)
	}
	req, err := http.NewRequest("POST", service.url, r.Body)
	req.URL = parsedURL
	if err != nil {
		log.Fatalf("Error in creating request to hit Spotify endpoint %v", err)
	}
	byteAuthorization := []byte(service.clientId + ":" + service.secret)
	basicAuthorization := "Basic " + b64.StdEncoding.EncodeToString(byteAuthorization)

	req.Header.Set("Authorization", basicAuthorization)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error getting response from the Spotify endpoint %v", err)
	}
	//log.Printf("%+v", req)
	//log.Printf("%+v", resp)
	tokenResponse := &dto.SpotifyResponse{}
	if err = utilities.FromJson(resp.Body, tokenResponse); err != nil {
		log.Fatalf("Error parsing response from the Spotify endpoint %v", err)
	}
	log.Printf("%v", tokenResponse.AccessToken)
	if tokenResponse.AccessToken == "" {
		tokenResponse.Success = false
		tokenResponse.Status = 400
	} else {
		tokenResponse.Status = 200
		tokenResponse.Success = true
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	//log.Printf("%v", tokenResponse)
	json.NewEncoder(rw).Encode(tokenResponse)

}
