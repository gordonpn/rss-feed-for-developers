# TODO

## Fetcher

- [x] Prepare SQL init scripts and schema
- [x] Insert the data into the database
- [x] Mount data.json from host machine to Docker container to allow externally changing the configuration
- [x] Schedule to run twice per day

## Backend

- [x] Endpoint that returns generated RSS 2.0 and return at `https://rss.gordon-pn.com/feed`
  - Grab anything that is from the last 6 months or last year
  - Return cache from Redis, cache valid for 24 hours
- [ ] Endpoint that returns feed items from the last 7 days
- [x] Cache the RSS using TTL tag
- [ ] Cache the RSS with Redis

## Frontend

- [ ] Explain what, why, how
- [ ] Show preview of the feed (items from the last week (7 days))
- [ ] Mount data.json from host machine to Docker container to show on the website which subreddits are being used for the feed at the moment

## CI/CD

- [ ] Drone CI pipeline
- [x] Docker compose configs
- [x] Deploy (prod/staging) and build scripts

## Proxy

- [ ] Need a NGINX or Caddy proxy
- [ ] Cache the RSS feed
