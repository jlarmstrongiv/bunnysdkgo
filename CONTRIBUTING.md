# Contributing

VS Code extensions

- `golang.go`

Create a `go.work` file with:

```go.work
go 1.22.4

replace github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client => ./bunnysdkgo/bunny_api_client
replace github.com/jlarmstrongiv/bunnysdkgo/edge_storage_api_client => ./bunnysdkgo/edge_storage_api_client
replace github.com/jlarmstrongiv/bunnysdkgo/logging_api_client => ./bunnysdkgo/logging_api_client
replace github.com/jlarmstrongiv/bunnysdkgo/stream_api_client => ./bunnysdkgo/stream_api_client

use .
```

<!-- GIT_TERMINAL_PROMPT https://stackoverflow.com/a/38237165 -->
<!-- update website https://stackoverflow.com/a/61974058 -->
<!-- request a manual website update at https://pkg.go.dev/github.com/jlarmstrongiv/bunnysdkgo@v0.0.1 -->

- Update dependencies `go get -u`
- `go mod tidy`
- `go build`
- `git add -A`
- `git commit -m "message"`
- `git push`
- `git tag v0.0.1`
- `git push origin v0.0.1`
- `GIT_TERMINAL_PROMPT=1 GOPROXY=proxy.golang.org go list -m github.com/jlarmstrongiv/bunnysdkgo@v0.0.1`
- `curl -X POST https://pkg.go.dev/fetch/github.com/jlarmstrongiv/bunnysdkgo@v0.0.1`
