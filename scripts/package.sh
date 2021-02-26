#!/usr/bin/env bash

go build -ldflags="-s -w" -o ./bin/detect ./cmd/detect/main.go && \
go build -ldflags="-s -w" -o ./bin/build ./cmd/build/main.go && \
pack buildpack package oscarnevarezleal/simple-packit-buildpack --config ./package.toml

pack build app-name --path sample-app --buildpack . --builder paketobuildpacks/builder:tiny --verbose
