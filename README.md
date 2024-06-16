# bunny-sdk

[![Go Reference](https://pkg.go.dev/badge/github.com/jlarmstrongiv/bunnysdkgo.svg)](https://pkg.go.dev/github.com/jlarmstrongiv/bunnysdkgo)

## Install

```shell
go get github.com/jlarmstrongiv/bunnysdkgo
```

## Documentation

Please read the [docs](https://bunny-launcher.com/bunny-sdk/go) for usage and examples.

## Contributing

Create a `go.work` file with:

```go.work
go 1.22.4

replace github.com/jlarmstrongiv/bunnysdkgo/generated => ./bunnysdkgo/generated

use .
```

<!-- GIT_TERMINAL_PROMPT https://stackoverflow.com/a/38237165 -->
<!-- update website https://stackoverflow.com/a/61974058 -->
<!-- request a manual website update at https://pkg.go.dev/github.com/jlarmstrongiv/bunnysdkgo@v0.0.1 -->

- `go mod tidy`
- `go build`
- `git add -A`
- `git commit -m "message"`
- `git push`
- `git tag v0.0.1`
- `git push origin v0.0.1`
- `GIT_TERMINAL_PROMPT=1 GOPROXY=proxy.golang.org go list -m github.com/jlarmstrongiv/bunnysdkgo@v0.0.1`
- `curl -X POST https://pkg.go.dev/fetch/github.com/jlarmstrongiv/bunnysdkgo@v0.0.1`
