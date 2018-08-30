package median

import "sync"

const bins = 16536
const procs = 4

type Interface interface {
	// Len is the number of elements in the collection
	Len() int
	//
	GetValue(i int) float64
	Split(a, b int) Interface
}

type Float64Slice []float64

func (t Float64Slice) Len() int {
	return len(t)
}

func (t Float64Slice) GetValue(i int) float64 {
	return t[i]
}

func (t Float64Slice) Split(a, b int) Interface {
	return t[a:b]
}

func Median(xs Interface, min float64, max float64) float64 {
	var median float64

	for {
		bias := 0
		hist := make([]int, bins, bins)

		delta := float64(max-min) / bins

		for i := 0; i < xs.Len(); i++ {
			val := xs.GetValue(i)
			if val > min && val < max {
				hist[int((val-min)/delta)]++
			}
		}

		ind, bias := getindex(hist, bias)
		min = min + delta*float64(ind)
		max = min + delta

		if hist[ind] == 1 {
			median = float64(findmedian(xs, min, max))
			break
		}
	}

	return median
}

func ParallelMedian(xs Interface, min float64, max float64) float64 {
	var median float64
	var mu sync.Mutex

	for {
		bias := 0
		acchist := make([]int, bins, bins)

		delta := float64(max-min) / bins

		var wg sync.WaitGroup
		wg.Add(procs)
		for p := 0; p < procs; p++ {
			go func(sxs Interface) {
				hist := make([]int, bins, bins)
				for i := 0; i < sxs.Len(); i++ {
					put(hist, min, max, sxs.GetValue(i), delta)
				}
				mu.Lock()
				for i := 0; i < bins; i++ {
					acchist[i] += hist[i]
				}
				mu.Unlock()
				wg.Done()
			}(xs.Split((xs.Len()*p)/procs, (xs.Len()*(p+1))/procs))
		}

		wg.Wait()

		ind, bias := getindex(acchist, bias)
		min = min + delta*float64(ind)
		max = min + delta

		if acchist[ind] == 1 {
			median = float64(findmedian(xs, min, max))
			break
		}
	}

	return median
}

func put(hist []int, min float64, max float64, val float64, delta float64) {
	if val < min || val > max {
		return
	}

	pos := int((val - min) / delta)
	hist[pos]++
}

func findmedian(values Interface, min float64, max float64) float64 {
	for i := 0; i < values.Len(); i++ {
		if values.GetValue(i) >= min && values.GetValue(i) <= max {
			return values.GetValue(i)
		}
	}

	return 0
}

func getindex(hist []int, bias int) (int, int) {
	left := 0
	right := 0
	var i int

	for k := 0; k < len(hist); k++ {
		right += hist[k]
	}

	for i = 0; i < len(hist); i++ {
		right -= hist[i]
		if i > 0 {
			left += hist[i-1]
		}
		if hist[i] > 0 {
			if bias+right-left <= hist[i] {
				bias = bias + right - left
				break
			}
		}

	}

	return i, bias
}
