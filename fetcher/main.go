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

var httpClient = &http.Client{Timeout: 10 * time.Second}

type DevToPost struct {
	Title       string
	Link        string
	Description string
	Published   time.Time
	Author      string
	ID          int
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn(fmt.Sprintf("Did not load .env file: %v", err))
	}
	log.SetLevel(log.DebugLevel)
}

func main() {
	_ = fetchDevToPosts()
	userAgent := "github.com/gordonpn/rss-feed-for-developers by /u/gordonpn"
	reqBody := url.Values{}
	reqBody.Set("grant_type", "password")
	reqBody.Set("username", os.Getenv("REDDIT_USERNAME"))
	reqBody.Set("password", os.Getenv("REDDIT_PASSWORD"))
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", strings.NewReader(reqBody.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(os.Getenv("REDDIT_ID"), os.Getenv("REDDIT_SECRET"))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", respData)
	var respMap map[string]interface{}
	if err := json.Unmarshal(respData, &respMap); err != nil {
		log.Fatal(err)
	}
	accessToken, _ := respMap["access_token"].(string)
	requ, err := http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
	if err != nil {
		log.Fatal(err)
	}
	requ.Header.Set("Authorization", fmt.Sprintf("bearer %s", accessToken))
	requ.Header.Set("User-Agent", userAgent)
	respo, err := httpClient.Do(requ)
	if err != nil {
		log.Fatal(err)
	}
	defer respo.Body.Close()
	respoData, err := ioutil.ReadAll(respo.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", respoData)
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
		id, _ := post["id"].(int)
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
	s, _ := json.MarshalIndent(devPosts, "", "\t")
	fmt.Printf("%s\n", s)
	fmt.Printf("Length: %d\n", len(devPosts))
	return
}
