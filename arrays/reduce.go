package arrays

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
