package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn(fmt.Sprintf("Did not load .env file: %v", err))
	}
	log.SetLevel(log.DebugLevel)
}

func main() {
	var resMap []map[string]interface{}
	err := getJson("https://dev.to/api/articles?top=30", &resMap)
	if err != nil {
		log.Fatal(err)
	}
	s, _ := json.MarshalIndent(resMap, "", "\t")
	fmt.Printf("%s\n", s)
}

func getJson(url string, target interface{}) error {
	res, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
