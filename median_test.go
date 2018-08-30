package median

import "testing"
import "math/rand"
import "strconv"
import "github.com/montanaflynn/stats"

type testCase struct {
	length int
	data   []float64
	expect float64
}

var testCases = []testCase{{1024 + 1, []float64{}, 1.0}, {2048 + 1, []float64{}, 3.0}, {4096 + 1, []float64{}, 5.0}, {16777216 + 1, []float64{}, 5.0}}

func init() {
	for i := 0; i < len(testCases); i++ {
		testCases[i].data = generateRandom(testCases[i].length)
		testCases[i].expect, _ = stats.Median(testCases[i].data)
	}
}

func generateRandom(length int) []float64 {
	r := make([]float64, length, length)

	for i := 0; i < length; i++ {
		r[i] = rand.Float64()
	}

	return r
}

func TestMedian(t *testing.T) {
	for i := 0; i < len(testCases)-1; i++ {
		m := ParallelMedian(Float64Slice(testCases[i].data), 0.0, 1.0)
		if m != testCases[i].expect {
			t.Errorf("Median was incorrect, got: %v, want: %v.", m, testCases[i].expect)
		}
	}
}

func BenchmarkParallelMedian(b *testing.B) {
	for i := 2; i < testCases[len(testCases)-1].length; i *= 2 {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				ParallelMedian(Float64Slice(testCases[len(testCases)-1].data[0:i]), 0.0, 1.0)
			}
		})

	}
}

func BenchmarkMedianSorting(b *testing.B) {
	for i := 2; i < testCases[len(testCases)-1].length; i *= 2 {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				stats.Median(testCases[len(testCases)-1].data[0:i])
			}
		})
	}
}
