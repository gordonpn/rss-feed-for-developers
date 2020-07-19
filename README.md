# RSS Feed for Developers

[![Project Status: WIP – Initial development is in progress](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)
[![Go Report Card](https://goreportcard.com/badge/github.com/gordonpn/rss-feed-for-developers)](https://goreportcard.com/report/github.com/gordonpn/rss-feed-for-developers)

## Description & Motivation

I like to consume Reddit, but ideally in a low-volume amount. It is way too easy to open Reddit and start going down the rabbit hole.

I like Dev.to as well, but let's be honest, there's a lot of badly written and less interesting articles there.

### Problem 1

One could easily turn any subreddit into an RSS feed by appending `/.rss` to the URL. For whatever reason, some of the feeds I'd like to follow were not updated in the last year or more!

### Problem 2

The Dev.to RSS feed at [`https://dev.to/feed`](https://dev.to/feed) returns way too much content. Really cluttering one's RSS reader.

### My Solution

Generate my own RSS feed!

## How It Works

Deployed on a bare metal home server, the `fetcher` service gets the top 2 posts from each of subreddit in a list, and the top 30 posts from the last 30 days from Dev.to, both twice per day.

The `backend` service will return (WIP) the generated feed as a valid RSS that anybody can subscribe through their preferred RSS reader, mine is Feedly.
