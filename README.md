### goactivesupport

[![Maintainability](https://api.codeclimate.com/v1/badges/1ea328c26d422c1bc415/maintainability)](https://codeclimate.com/github/ipoval/goactivesupport/maintainability)

[goreportcard.com](https://goreportcard.com/report/github.com/ipoval/goactivesupport)

#### REST package
```go
import (
  gasrest "github.com/ipoval/goactivesupport/rest"
)

gasrest.UserAgent.Chrome()
gasrest.UserAgent.Random()
```

Tests
```shell script
make help
setup                          Install all the build and lint dependencies
dep                            Run dep ensure and prune
test                           Run all the tests
cover                          Run all the tests and opens the coverage report
fmt                            Run goimports on all go files
lint                           Run all the linters
ci                             Run all the tests and code checks
tagver                         Tag the branch and push new git tag to origin
build                          Build a version
clean                          Remove temporary files
help                           make help

```

Benchmark
```shell script
go test -v -run="none" -bench=. -benchtime="3s" -benchmem
goos: darwin
goarch: amd64
BenchmarkNumberToStrFast-4      100000000               33.6 ns/op             5 B/op          1 allocs/op
BenchmarkSprintf-4              44198298                80.9 ns/op             5 B/op          1 allocs/op
```
