package main

import (
	"fmt"
	"time"

	types "github.com/gordonpn/rss-feed-for-developers/fetcher/pkg"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func fetchDevToPosts() (devPosts []types.Post) {
	log.Info("Fetching Dev.to posts")
	var resp []map[string]interface{}
	err := getJSON("https://dev.to/api/articles?top=30", &resp, nil)
	checkAndPanic("Error with parsing JSON", err)
	for _, post := range resp {
		title, _ := post["title"].(string)
		title = cleanString(title)
		link, _ := post["canonical_url"].(string)
		description, _ := post["description"].(string)
		description = cleanString(description)
		publishedString, _ := post["published_at"].(string)
		published, _ := time.Parse(time.RFC3339, publishedString)
		author, _ := post["user"].(map[string]interface{})["name"].(string)
		author = cleanString(author)
		idFloat, _ := post["id"].(float64)
		id := fmt.Sprintf("%.0f", idFloat)
		aPost := types.Post{
			Title:       title,
			Link:        link,
			Description: description,
			Published:   published,
			Author:      author,
			ID:          id,
		}
		devPosts = append(devPosts, aPost)
	}
	log.Info("Fetching Dev.to posts: done")
	return
}
