package fibona

func Fbb(n int) {
	seq := make([]int, n)
	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	LastValue = seq[n-1]
}

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
