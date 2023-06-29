package fibo

var LastValue int

func CalculateFibonacci(n int) {
	first := 0
	second := 1

	for i := 2; i <= n; i++ {
		next := first + second
		first = second
		second = next
	}

	LastValue = second
}

func Operation(value int) int {
	return value
}
