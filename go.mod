module github.com/gordonpn/rss-feed-for-developers

go 1.14

replace github.com/gordonpn/rss-feed-for-developers/fetcher => ./fetcher

replace github.com/gordonpn/rss-feed-for-developers/backend => ./backend

require github.com/gordonpn/rss-feed-for-developers/fetcher v0.0.0-00010101000000-000000000000 // indirect
