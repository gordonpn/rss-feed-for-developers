package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/feeds"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "postgres"
	}
	port := 5432
	pgURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connectDB(a, pgURI)
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) handleFeed() http.HandlerFunc {
	log.Debug("Feed endpoint registered")
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		feed := &feeds.Feed{
			Title:       "jmoiron.net blog",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
			Description: "discussion about tech, footie, photos",
			Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
		}

		feed.Items = []*feeds.Item{
			{
				Title:       "Limiting Concurrency in Go",
				Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
				Description: "A discussion on controlled parallelism in golang",
				Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
				Created:     now,
				Id:          "http://jmoiron.net/blog/limiting-concurrency-in-go",
			},
			{
				Title:       "Logic-less Template Redux",
				Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
				Description: "More thoughts on logicless templates",
				Created:     now,
				Id:          "http://jmoiron.net/blog/logicless-template-redux",
			},
			{
				Title:       "Idiomatic Code Reuse in Go",
				Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
				Description: "How to use interfaces <em>effectively</em>",
				Created:     now,
				Id:          "http://jmoiron.net/blog/idiomatic-code-reuse-in-go",
			},
		}

		rss, err := feed.ToRss()
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/rss+xml")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(rss))
		if err != nil {
			log.Warn(err)
		}
	}
}

func (a *App) handleHealthCheck() http.HandlerFunc {
	log.Debug("Healthcheck API endpoint registered")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Healthcheck pinged")
		if err := a.DB.Ping(); err != nil {
			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "not ok"})
			log.Error(err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "ok"})
	}
}

func (a *App) handlePosts() http.HandlerFunc {
	log.Debug("Posts API endpoint registered")
	return func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, http.StatusInternalServerError, "not yet implemented")
	}
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/feed", a.handleFeed()).Methods("GET")
	apiRoute := a.Router.PathPrefix("/api/v1").Subrouter()
	apiRoute.HandleFunc("/healthcheck", a.handleHealthCheck()).Methods("GET")
	apiRoute.HandleFunc("/posts", a.handlePosts()).Methods("GET")
}
