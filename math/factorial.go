package math

func Factorial(a int) int {
	if a < 0 {
		panic("expected non-negative integer")
	}
	res := a
	for i := a - 1; i > 0; i-- {
		res *= i
	}
	return res
}
