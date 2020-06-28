### goactivesupport

[![Maintainability](https://api.codeclimate.com/v1/badges/1ea328c26d422c1bc415/maintainability)](https://codeclimate.com/github/ipoval/goactivesupport/maintainability)


#### REST package
```go

```

Tests
```shell script
make test
```

Benchmark
```shell script
go test -v -run="none" -bench=. -benchtime="3s" -benchmem
goos: darwin
goarch: amd64
BenchmarkNumberToStrFast-4      100000000               33.6 ns/op             5 B/op          1 allocs/op
BenchmarkSprintf-4              44198298                80.9 ns/op             5 B/op          1 allocs/op
```
