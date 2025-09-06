package distributions

import "math"

func Poisson(lambda float64, x []float64) []float64 {
	res := make([]float64, len(x))
	for i, v := range x {
		if v < 0 || math.Abs(v-math.Round(v)) > 1e-9 {
			res[i] = 0
			continue
		}
		res[i] = math.Pow(lambda, v) / factorial(v) * math.Exp(-lambda)
	}
	return res
}
