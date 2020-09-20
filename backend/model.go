package main

import (
	types "github.com/gordonpn/rss-feed-for-developers/fetcher/pkg"
	. "github.com/gorilla/feeds" //nolint:golint
	"time"
)

func (a *App) getFeed() (string, error) {
	now := time.Now()
	feed := &Feed{
		Title:       "Feed for Developers",
		Link:        &Link{Href: "https://rss.gordon-pn.com"},
		Description: "Top articles and posts from Dev.to and Reddit",
		Author:      &Author{Name: "gordonpn", Email: "rss@gordon-pn.com"},
		Created:     now,
	}

	feed.Items = []*Item{}

	query := `SELECT * FROM posts WHERE published > CURRENT_TIMESTAMP - INTERVAL '91 day'`
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
		thisItem := &Item{
			Title:       thisPost.Title,
			Link:        &Link{Href: thisPost.Link},
			Author:      &Author{Name: thisPost.Author},
			Description: thisPost.Description,
			Id:          thisPost.Link,
			Created:     thisPost.Published,
		}
		feed.Items = append(feed.Items, thisItem)
	}
	rssFeed := (&Rss{Feed: feed}).RssFeed()
	rssFeed.Ttl = 720
	rss, err := ToXML(rssFeed)
	return rss, err
}
