package slice

func Sum(numbers []int) int {
	var sum int

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numSlices ...[]int) []int {
	result := []int{}

	for _, nums := range numSlices {
		sum := Sum(nums)
		result = append(result, sum)
	}
	return result
}

func SumTail(numSlices ...[]int) []int {
	result := []int{}

	for _, nums := range numSlices {
		if len(nums) == 0 {
			result = append(result, 0)
			continue
		}

		tails := nums[1:]

		result = append(result, Sum(tails))
	}

	return result
}

// 也可以這樣寫
// func SumAll(numSlices ...[]int) []int {
// 	result := make([]int, len(numSlices))

// 	for i, nums := range numSlices {
// 		sum := Sum(nums)
// 		result[i] = sum
// 	}
// 	return result
// }
