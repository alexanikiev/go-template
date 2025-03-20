# go-template

Go template code

## Web API

Default: net/http
Options: Gin, Echo, Fiber

Tutorial: https://go.dev/doc/tutorial/web-service-gin

### Build

`go build ./cmd/...`

### Test

`go test ./pkg/solver/`

### Run

`go run ./cmd/api/main.go`

## Running GitHub actions locally

https://github.com/nektos/act

`act -W .github/workflows/ci.yaml -j lint-format-test`
