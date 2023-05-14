package math

func Addition(x, y float64) float64 {
	return x + y
}

func Multiplication(x, y float64) float64 {
	return x * y
}

func Subtraction(x, y float64) float64 {
	return x - y
}

func Division(x, y float64) float64 {
       if y != 0 {
		return x / y
	}
	return 0
}
