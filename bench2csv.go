package bench2csv

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// benchmakrMatcher matches a benchmark output line.
// See https://regex101.com/r/Uv4LNN/latest
var benchmarkMatcher = regexp.MustCompile(
	`^Benchmark` + // "Benchmark" prefix
		`(?P<name>[^-\s]+)` + // Name
		`(?:-(?P<parallelism>\d+))?` + // Optional parallelism (set with -cpu flag)
		`\s+` +
		`(?P<ops>\d+)` + // Operations run
		`\s+` +
		`(?P<duration>\d+(?:\.\d+)?)` + // Duration for each operation
		`\sns/op` + // Duration unit suffix
		`$`)

// Process benchmark output from in, write CSV to csvOut, and pipe benchmark output to errOut.
func Process(in io.Reader, csvOut, errOut io.Writer) error {
	s := bufio.NewScanner(in)

	if _, err := fmt.Fprintln(csvOut, "name,parallelism,ops,duration"); err != nil {
		return err
	}

	for s.Scan() {
		line := s.Text()
		if _, err := fmt.Fprintln(errOut, line); err != nil {
			return err
		}

		if !benchmarkMatcher.MatchString(line) {
			continue
		}

		submatches := benchmarkMatcher.FindStringSubmatch(line)
		submatches = submatches[1:]

		name := submatches[0]
		parallelism := submatches[1]
		ops := submatches[2]
		duration := submatches[3]

		if parallelism == "" {
			parallelism = "1"
		}

		if _, err := fmt.Fprintln(csvOut, strings.Join([]string{name, parallelism, ops, duration}, ",")); err != nil {
			return err
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}
