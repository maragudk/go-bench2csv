# go-bench2csv

A small CLI to parse the output of `go test -bench` and output to CSV.

It passes the `go test` output verbatim on STDERR and the CSV output to STDOUT.

## Usage

```shell
$ go install ./cmd/bench2csv
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

Made in 🇩🇰 by [maragu](https://www.maragu.dk/), maker of [online Go courses](https://www.golang.dk/).
