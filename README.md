# Simple buildpack using PackIt

## Compile and pack this buildpack
```bash
go build -ldflags="-s -w" -o ./bin/build ./cmd/build/main.go && \
pack buildpack package oscarnevarezleal/simple-packit-buildpack --config ./package.toml
```

## Test it with the included sample-app
```bash
pack build app-name --path sample-app --buildpack . --builder paketobuildpacks/builder:tiny --verbose
```