package bench2csv_test

import (
	"embed"
	"io"
	"strings"
	"testing"

	"github.com/maragudk/go-bench2csv"
)

//go:embed testdata
var testdata embed.FS

func TestProcess(t *testing.T) {
	t.Run("outputs CSV with name,parallelism,operations,duration", func(t *testing.T) {
		var csv strings.Builder

		err := bench2csv.Process(strings.NewReader(readFile(t, "input1.txt")), &csv, io.Discard, bench2csv.Default)
		noErr(t, err)

		if csv.String() != readFile(t, "output1.csv") {
			t.Fatal("Unexpected output:", csv.String())
		}
	})

	t.Run("outputs CSV with frequency if set", func(t *testing.T) {
		var csv strings.Builder

		err := bench2csv.Process(strings.NewReader(readFile(t, "input1.txt")), &csv, io.Discard,
			bench2csv.Default|bench2csv.Frequency)
		noErr(t, err)

		if csv.String() != readFile(t, "output2.csv") {
			t.Fatal("Unexpected output:", csv.String())
		}
	})

	t.Run("outputs CSV with benchmem statistics if set", func(t *testing.T) {
		var csv strings.Builder

		err := bench2csv.Process(strings.NewReader(readFile(t, "input2.txt")), &csv, io.Discard,
			bench2csv.Default|bench2csv.Mem)
		noErr(t, err)

		if csv.String() != readFile(t, "output3.csv") {
			t.Fatal("Unexpected output:", csv.String())
		}
	})

	t.Run("pipes input to output", func(t *testing.T) {
		input := readFile(t, "input1.txt")

		var out strings.Builder

		err := bench2csv.Process(strings.NewReader(input), io.Discard, &out, bench2csv.Default)
		noErr(t, err)

		if out.String() != input {
			t.Fatal("Unexpected output:", out.String())
		}
	})
}

func BenchmarkProcess(b *testing.B) {
	b.Run("in parallel just for fun", func(b *testing.B) {
		input := readFile(b, "input1.txt")

		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				err := bench2csv.Process(strings.NewReader(input), io.Discard, io.Discard, 0)
				noErr(b, err)
			}
		})
	})
}

func noErr(tb testing.TB, err error) {
	if err != nil {
		tb.Fatal(err)
	}
}

func readFile(tb testing.TB, name string) string {
	d, err := testdata.ReadFile("testdata/" + name)
	if err != nil {
		tb.Fatal(err)
	}
	return string(d)
}
