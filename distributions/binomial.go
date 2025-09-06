package distributions

import "math"

func factorial(n float64) float64 {
	int_n := int(n)
	res := 1.0

	if int_n < 0 {
		panic("factorial: negative input not allowed")
	}
	if int_n > 170 {
		panic("factorial: input too large, would cause overflow")
	}

	if int_n == 0 || int_n == 1 {
		return res
	}

	for i := 2; i <= int(n); i++ {
		res *= float64(i)
	}

	return res
}

func combination(n, k int) float64 {
	if k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}

	return factorial(float64(n)) / (factorial(float64(k)) * factorial(float64(n-k)))
}

func Binomial(n int, p float64, x []float64) []float64 {
	res := make([]float64, len(x))
	for i, v := range x {
		if v < 0 || v > float64(n) || math.Abs(v-math.Round(v)) > 1e-9 {
			res[i] = 0
			continue
		}
		res[i] = combination(n, int(v)) * math.Pow(p, v) * math.Pow(1-p, float64(n)-v)
	}
	return res
}
