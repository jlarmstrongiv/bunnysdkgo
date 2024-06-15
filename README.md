# [`bunny-sdk`](https://bunny-launcher.com/bunny-sdk/go)

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

- `go mod tidy`
- `go build`
- `git tag v0.0.1`
- `git push origin v0.0.1`
- `GOPROXY=proxy.golang.org go list -m github.com/jlarmstrongiv/bunnysdkgo@v0.0.1`
