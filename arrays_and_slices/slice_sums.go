package list

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SliceSums(slices ...[]int) []int {
	sliceCounts := len(slices)
	sums := make([]int, sliceCounts)

	for i, slice := range slices {
		sums[i] = Sum(slice)
	}

	return sums
}
