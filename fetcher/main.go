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

type DevToPost struct {
	Title       string
	Link        string
	Description string
	Published   time.Time
	Author      string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn(fmt.Sprintf("Did not load .env file: %v", err))
	}
	log.SetLevel(log.DebugLevel)
}

func main() {
	_ = fetchDevToPosts()
}

func getJSON(url string, target interface{}) error {
	res, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

func fetchDevToPosts() (devPosts []DevToPost) {
	var resMap []map[string]interface{}
	err := getJSON("https://dev.to/api/articles?top=30", &resMap)
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range resMap {
		title, _ := post["title"].(string)
		link, _ := post["canonical_url"].(string)
		description, _ := post["description"].(string)
		publishedString, _ := post["published_at"].(string)
		published, _ := time.Parse(time.RFC3339, publishedString)
		author, _ := post["user"].(map[string]interface{})["name"].(string)
		aPost := DevToPost{Title: title, Link: link, Description: description, Published: published, Author: author}
		devPosts = append(devPosts, aPost)
	}
	s, _ := json.MarshalIndent(devPosts, "", "\t")
	fmt.Printf("%s\n", s)
	fmt.Printf("Length: %d", len(devPosts))
	return
}
