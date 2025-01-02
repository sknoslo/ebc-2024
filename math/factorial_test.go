package math

import "testing"

func TestFactorial(t *testing.T) {
	testCase := func (a, want int) {
		actual := Factorial(a)
		if actual != want {
			t.Errorf("Factorial(%d) = %d, want %d", a, actual, want)
		}
	}

	testCase(0, 0)
	testCase(1, 1)
	testCase(2, 2)
	testCase(3, 6)
	testCase(10, 3_628_800)
}
