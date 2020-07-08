package types

import "time"

type Post struct {
	Title       string
	Link        string
	Description string
	Published   time.Time
	Author      string
	ID          string
}
