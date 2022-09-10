package math

func Add(x, y int) int {

	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Div(x, y int) float64 {
	if y == 0 {
		return float64(0)
	}
	return float64(x / y)
}

func Multi(x, y int) int {
	return x * y
}
