#!/bin/sh
echo "$GITHUB_TOKEN" | docker login ghcr.io -u gordonpn --password-stdin
docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
docker buildx rm builder || true
docker buildx create --name builder --driver docker-container --use
docker buildx inspect --bootstrap
DOCKER_TAG=${DOCKER_TAG:-latest}
cd /drone/src/backend || exit 1
docker buildx build -t ghcr.io/gordonpn/rss-feed_backend:"$DOCKER_TAG" --platform linux/amd64,linux/arm64 --push .
cd /drone/src/fetcher || exit 1
docker buildx build -t ghcr.io/gordonpn/rss-feed_fetcher:"$DOCKER_TAG" --platform linux/amd64,linux/arm64 --push .
