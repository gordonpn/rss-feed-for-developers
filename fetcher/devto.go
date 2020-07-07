package main

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type DevToPost struct {
	Title       string
	Link        string
	Description string
	Published   time.Time
	Author      string
	ID          int
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func fetchDevToPosts() (devPosts []DevToPost) {
	log.Info("Fetching Dev.to posts")
	var resMap []map[string]interface{}
	err := getJSON("https://dev.to/api/articles?top=30", &resMap, nil)
	checkAndPanic(err)
	for _, post := range resMap {
		title, _ := post["title"].(string)
		link, _ := post["canonical_url"].(string)
		description, _ := post["description"].(string)
		publishedString, _ := post["published_at"].(string)
		published, _ := time.Parse(time.RFC3339, publishedString)
		author, _ := post["user"].(map[string]interface{})["name"].(string)
		idFloat, _ := post["id"].(float64)
		id := int(idFloat)
		aPost := DevToPost{
			Title:       title,
			Link:        link,
			Description: description,
			Published:   published,
			Author:      author,
			ID:          id,
		}
		devPosts = append(devPosts, aPost)
	}
	//s, _ := json.MarshalIndent(devPosts, "", "\t")
	//fmt.Printf("%s\n", s)
	//fmt.Printf("Length: %d\n", len(devPosts))
	log.Info("Done fetching Dev.to posts")
	return
}
