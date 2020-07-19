package main

import (
	types "github.com/gordonpn/rss-feed-for-developers/fetcher/pkg"
	"github.com/gorilla/feeds"
	"time"
)

func (a *App) getFeed() (string, error) {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Feed for Developer",
		Link:        &feeds.Link{Href: "https://rss.gordon-pn.com"},
		Description: "Top articles and posts from Dev.to and Reddit",
		Author:      &feeds.Author{Name: "gordonpn", Email: "gordon.pn6@gmail.com"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{}

	query := `SELECT * FROM posts WHERE published > CURRENT_TIMESTAMP - INTERVAL '183 day'`
	rows, err := a.DB.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		thisPost := types.Post{}
		if err = rows.Scan(
			&thisPost.Title,
			&thisPost.Link,
			&thisPost.Description,
			&thisPost.Published,
			&thisPost.Author,
			&thisPost.ID,
		); err != nil {
			return "", err
		}
		thisItem := &feeds.Item{
			Title:       thisPost.Title,
			Link:        &feeds.Link{Href: thisPost.Link},
			Author:      &feeds.Author{Name: thisPost.Author},
			Description: thisPost.Description,
			Id:          thisPost.Link,
			Created:     thisPost.Published,
		}
		feed.Items = append(feed.Items, thisItem)
	}
	rss, err := feed.ToRss()
	return rss, err
}
