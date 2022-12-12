package main

import (
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
	return bench2csv.Process(os.Stdin, os.Stdout, os.Stderr)
}
