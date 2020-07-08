package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var userAgent = "github.com/gordonpn/rss-feed-for-developers by /u/gordonpn"

type RedditPost struct {
	Title     string
	Link      string
	Published time.Time
	Author    string
	ID        string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn(fmt.Sprintf("Did not load .env file: %v", err))
	}
	log.SetLevel(log.DebugLevel)
}

func authenticate() (accessToken string) {
	log.Info("Authenticating Reddit API")
	reqBody := url.Values{}
	reqBody.Set("grant_type", "password")
	reqBody.Set("username", os.Getenv("REDDIT_USERNAME"))
	reqBody.Set("password", os.Getenv("REDDIT_PASSWORD"))
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(reqBody.Encode()))
	checkAndPanic(err)
	req.SetBasicAuth(os.Getenv("REDDIT_ID"), os.Getenv("REDDIT_SECRET"))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := httpClient.Do(req)
	checkAndPanic(err)
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	checkAndPanic(err)
	//fmt.Printf("%s\n", respData)
	var respMap map[string]interface{}
	err = json.Unmarshal(respData, &respMap)
	checkAndPanic(err)
	accessToken, _ = respMap["access_token"].(string)
	log.Info("Done authenticating Reddit API")
	return
}

func fetchRedditListings(subreddits []string) (redditPosts []RedditPost) {
	log.Info("Fetching top posts from Reddit")
	accessToken := authenticate()
	headers := map[string]string{
		"Authorization": fmt.Sprintf("bearer %s", accessToken),
		"User-Agent":    userAgent,
	}
	for _, subreddit := range subreddits {
		var resp map[string]interface{}
		redditURL := fmt.Sprintf("https://oauth.reddit.com/r/%s/top?limit=2&t=day", subreddit)
		log.Info(fmt.Sprintf("Processing: %s", redditURL))
		err := getJSON(redditURL, &resp, headers)
		checkAndPanic(err)
		data := resp["data"].(map[string]interface{})
		children := data["children"].([]interface{})
		for _, child := range children {
			aChild := child.(map[string]interface{})
			childData := aChild["data"].(map[string]interface{})
			title := childData["title"].(string)
			permalink := childData["permalink"].(string)
			author := childData["author"].(string)
			createdTime := childData["created_utc"].(float64)
			published := time.Unix(int64(createdTime), 0)
			id := childData["id"].(string)
			aPost := RedditPost{
				Title:     title,
				Link:      fmt.Sprintf("https://reddit.com%s", permalink),
				Author:    author,
				Published: published,
				ID:        id,
			}
			redditPosts = append(redditPosts, aPost)
		}
		log.Info(fmt.Sprintf("Done processing: %s", redditURL))
		//s, _ := json.MarshalIndent(redditPosts, "", "\t")
		//fmt.Printf("%s\n", s)
	}
	return
}

func getSubreddits() (subreddits []string) {
	log.Info("Opening data file")
	jsonFile, err := os.Open("data.json")
	checkAndPanic(err)
	log.Info("Successfully opened data file")
	defer jsonFile.Close()
	var data map[string]interface{}
	err = json.NewDecoder(jsonFile).Decode(&data)
	checkAndPanic(err)
	subs := data["subreddits"].([]interface{})
	for _, sub := range subs {
		subreddits = append(subreddits, sub.(string))
	}
	//fmt.Printf("%s\n", subreddits)
	return
}
