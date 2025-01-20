package math

func Lcm(a, b int) int {
	return (a * b) / Gcd(a, b)
}
