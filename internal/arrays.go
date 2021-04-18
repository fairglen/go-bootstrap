package internal

// https://dev.to/quii/learn-go-by-writing-tests-arrays-and-slices-ahm

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum+= number
	}
	return
}
