package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maragudk/go-bench2csv"
)

func main() {
	if err := start(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func start() error {
	format := bench2csv.Default

	freq := flag.Bool("freq", false, "Include frequency output")
	mem := flag.Bool("mem", false, "Include -benchmem output")

	flag.Parse()

	if *freq {
		format |= bench2csv.Frequency
	}

	if *mem {
		format |= bench2csv.Memory
	}

	return bench2csv.Process(os.Stdin, os.Stdout, os.Stderr, format)
}
