package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	types "github.com/gordonpn/rss-feed-for-developers/fetcher/pkg"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	connection *sql.DB
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn(fmt.Sprintf("Did not load .env file: %v", err))
	}
	log.SetLevel(log.DebugLevel)
}

func (database *Database) connect() error {
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		host = "postgres"
	}
	user := os.Getenv("POSTGRES_NONROOT_USER")
	password := os.Getenv("POSTGRES_NONROOT_PASSWORD")
	dbname := "app_database"
	port := 5432
	pgURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Info("Connecting to database...")
	var err error
	for i := 1; i < 6; i++ {
		if database.connection, err = sql.Open("postgres", pgURI); err != nil {
			log.Error("Error with opening connection with database")
			return err
		}

		if err = database.connection.Ping(); err == nil {
			break
		}
		retryWait := i * i
		message := fmt.Sprintf("Connection attempt %d unsuccessful, retrying in %d seconds...", i, retryWait)
		log.Infof(message)
		time.Sleep(time.Duration(retryWait) * time.Second)
	}
	if database.connection == nil {
		log.Fatal("Could not connect to database")
	}
	log.Info("Successfully connected to database")
	return nil
}

func (database *Database) insert(items []types.Post) error {
	log.Info("Inserting posts into database")
	for _, item := range items {
		query := `INSERT INTO posts (title, link, summary, published, author, id)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (id) DO NOTHING`
		_, err := database.connection.Exec(query, item.Title, item.Link, item.Description, item.Published, item.Author, item.ID)
		if err != nil {
			return err
		}
	}
	log.Info("Inserting posts into database: done")
	return nil
}
