# RSS Feed for Developers

[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)
[![Go Report Card](https://goreportcard.com/badge/github.com/gordonpn/rss-feed-for-developers)](https://goreportcard.com/report/github.com/gordonpn/rss-feed-for-developers)
[![Build Status](https://drone.gordon-pn.com/api/badges/gordonpn/rss-feed-for-developers/status.svg)](https://drone.gordon-pn.com/gordonpn/rss-feed-for-developers)
![Healthchecks.io](https://healthchecks.io/badge/a75d9381-9b57-4452-9b52-800917001d96/1AJW2wU7-2.svg)

## Description & Motivation

I like to consume Reddit, but ideally in a low-volume amount. It is way too easy to open Reddit and start going down the rabbit hole.

I like Dev.to as well, but it requires too much sifting through.

### Problem 1

One could easily turn any subreddit into an RSS feed by appending `/.rss` to the URL. For whatever reason, some of the feeds I'd like to follow were not updated in the last year or more!

### Problem 2

The Dev.to RSS feed at [`https://dev.to/feed`](https://dev.to/feed) returns way too much content. Really cluttering one's RSS reader.

### My Solution

Generate my own RSS feed!

## How It Works

Deployed on a bare metal home server, the `fetcher` service gets the top 30 posts (from the past month) from each of the subreddits in a config file, and the top 30 posts from the last 30 days from Dev.to, both twice per day.

The `backend` service will return the generated feed as a valid RSS that anybody can subscribe through their preferred RSS reader. E.g., Feedly.

In the future, I'd like to make it possible for visitors to create a their own personalized feeds (with a unique URL to subscribe to). _Keyword: future_.
