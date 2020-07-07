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
)

var userAgent = "github.com/gordonpn/rss-feed-for-developers by /u/gordonpn"

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

func fetchRedditListings() {
	log.Info("Fetching top posts from Reddit")
	accessToken := authenticate()
	headers := map[string]string{
		"Authorization": fmt.Sprintf("bearer %s", accessToken),
		"User-Agent":    userAgent,
	}
	var resp map[string]interface{}
	redditURL := "https://oauth.reddit.com/api/v1/me"
	log.Info(fmt.Sprintf("Processing: %s", redditURL))
	err := getJSON(redditURL, &resp, headers)
	checkAndPanic(err)
	log.Info(fmt.Sprintf("Done processing: %s", redditURL))
	//s, _ := json.MarshalIndent(resp, "", "\t")
	//fmt.Printf("%s\n", s)
}
