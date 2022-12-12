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
	t.Run("outputs CSV with name,parallelism,ops,duration", func(t *testing.T) {
		var csv strings.Builder

		err := bench2csv.Process(strings.NewReader(readFile(t, "input1.txt")), &csv, io.Discard)
		noErr(t, err)

		if csv.String() != readFile(t, "output1.csv") {
			t.Fatal("Unexpected output:", csv.String())
		}
	})

	t.Run("pipes input to output", func(t *testing.T) {
		input := readFile(t, "input1.txt")

		var out strings.Builder

		err := bench2csv.Process(strings.NewReader(input), io.Discard, &out)
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
				err := bench2csv.Process(strings.NewReader(input), io.Discard, io.Discard)
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
