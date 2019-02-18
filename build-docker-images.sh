#!/bin/bash

docker build . -f Dockerfile -t tyhjataulu/go-edge-api:arm
docker build . -f Dockerfile.x86 -t tyhjataulu/go-edge-api:x86

docker push tyhjataulu/go-edge-api:arm
docker push tyhjataulu/go-edge-api:x86

docker manifest create --amend tyhjataulu/go-edge-api:latest \
tyhjataulu/go-edge-api:arm \
tyhjataulu/go-edge-api:x86

docker manifest push tyhjataulu/go-edge-api:latest