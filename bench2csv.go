package bench2csv

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

const (
	Name = 1 << iota
	Parallelism
	Operations
	Duration
	Frequency
	BytesPerOp
	AllocsPerOp
	Default = Name | Parallelism | Operations | Duration
	Memory  = BytesPerOp | AllocsPerOp
)

// benchmarkMatcher matches a benchmark output line.
// See https://regex101.com/r/EEQMWQ/latest
var benchmarkMatcher = regexp.MustCompile(
	`^Benchmark` + // "Benchmark" prefix
		`(?P<name>[^-\s]+)` + // Name
		`(?:-(?P<parallelism>\d+))?` + // Optional parallelism (set with -cpu flag)
		`\s+` +
		`(?P<ops>\d+)` + // Operations run
		`\s+` +
		`(?P<duration>\d+(?:\.\d+)?)` + // Duration for each operation
		`\sns/op` + // Duration unit suffix

		// Optionally, with -benchmem...
		`(\s+` +
		`(?P<bytesPerOp>\d+)` + // Bytes per operation
		`\sB/op\s+` + // Bytes per operation unit suffix
		`(?P<allocsPerOp>\d+)` + // Allocs per operation
		`\sallocs/op` + // Allocs per operation unit suffix
		`)?` +

		// The end
		`$`)

// Process benchmark output from in, write CSV to csvOut, and pipe benchmark output to errOut.
func Process(in io.Reader, csvOut, errOut io.Writer, format int) error {
	s := bufio.NewScanner(in)

	var header []string
	if format&Name != 0 {
		header = append(header, "name")
	}
	if format&Parallelism != 0 {
		header = append(header, "parallelism")
	}
	if format&Operations != 0 {
		header = append(header, "operations")
	}
	if format&Duration != 0 {
		header = append(header, "duration")
	}
	if format&Frequency != 0 {
		header = append(header, "frequency")
	}
	if format&BytesPerOp != 0 {
		header = append(header, "bytes_per_op")
	}
	if format&AllocsPerOp != 0 {
		header = append(header, "allocs_per_op")
	}

	if _, err := fmt.Fprintln(csvOut, strings.Join(header, ",")); err != nil {
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
		operations := submatches[2]
		duration := submatches[3]

		bytesPerOp := "0"
		allocsPerOp := "0"
		if len(submatches) > 6 {
			bytesPerOp = submatches[5]
			allocsPerOp = submatches[6]
		}

		if parallelism == "" {
			parallelism = "1"
		}

		durationAsFloat, err := strconv.ParseFloat(duration, 64)
		if err != nil {
			return err
		}
		frequency := strconv.FormatFloat(1e9/durationAsFloat, 'f', -1, 64)

		var values []string
		if format&Name != 0 {
			values = append(values, name)
		}
		if format&Parallelism != 0 {
			values = append(values, parallelism)
		}
		if format&Operations != 0 {
			values = append(values, operations)
		}
		if format&Duration != 0 {
			values = append(values, duration)
		}
		if format&Frequency != 0 {
			values = append(values, frequency)
		}
		if format&BytesPerOp != 0 {
			values = append(values, bytesPerOp)
		}
		if format&AllocsPerOp != 0 {
			values = append(values, allocsPerOp)
		}

		if _, err := fmt.Fprintln(csvOut, strings.Join(values, ",")); err != nil {
			return err
		}
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}
