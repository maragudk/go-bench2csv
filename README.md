# go-bench2csv

[![GoDoc](https://pkg.go.dev/badge/github.com/maragudk/go-bench2csv)](https://pkg.go.dev/github.com/maragudk/go-bench2csv)
[![Go](https://github.com/maragudk/go-bench2csv/actions/workflows/ci.yml/badge.svg)](https://github.com/maragudk/go-bench2csv/actions/workflows/ci.yml)

A small CLI to parse the output of `go test -bench` and output to CSV.

It passes the `go test` output verbatim on STDERR and the CSV output to STDOUT.

## Usage

![demo.gif](docs%2Fdemo.gif)

```shell
$ go install github.com/maragudk/go-bench2csv/cmd/bench2csv@latest
$ go test -cpu 1,2,4,8 -bench . | bench2csv >benchmark.csv
goos: darwin
goarch: arm64
pkg: github.com/maragudk/go-bench2csv
BenchmarkProcess/in_parallel_just_for_fun           	    4314	    276820 ns/op
BenchmarkProcess/in_parallel_just_for_fun-2         	    8292	    145504 ns/op
BenchmarkProcess/in_parallel_just_for_fun-4         	   15826	     75832 ns/op
BenchmarkProcess/in_parallel_just_for_fun-8         	   19111	     73122 ns/op
PASS
ok  	github.com/maragudk/go-bench2csv	7.324s
```

Also works with `go test -benchmem`:

```shell
$ go test -cpu 1,2,4,8 -bench . -benchmem | bench2csv -mem >benchmark.csv
goos: darwin
goarch: arm64
pkg: github.com/maragudk/go-bench2csv
BenchmarkProcess/in_parallel_just_for_fun           	    4106	    292497 ns/op	   53892 B/op	     738 allocs/op
BenchmarkProcess/in_parallel_just_for_fun-2         	    7929	    151227 ns/op	   53897 B/op	     738 allocs/op
BenchmarkProcess/in_parallel_just_for_fun-4         	   15013	     79910 ns/op	   53909 B/op	     738 allocs/op
BenchmarkProcess/in_parallel_just_for_fun-8         	   18214	     66196 ns/op	   53941 B/op	     738 allocs/op
PASS
ok  	github.com/maragudk/go-bench2csv	7.402s
```

Made in 🇩🇰 by [maragu](https://www.maragu.dk/), maker of [online Go courses](https://www.golang.dk/).
