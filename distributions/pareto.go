package distributions

import "math"

func Pareto(x0, alpha float64, x []float64) []float64 {
	res := make([]float64, len(x))

	for i, v := range x {
		if v < x0 {
			res[i] = 0
		} else {
			res[i] = (alpha * math.Pow(x0, alpha)) / math.Pow(v, alpha+1)
		}
	}

	return res
}
