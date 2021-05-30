package arrays

// The size of an arr in go is encoded in its type.
// Meaning, if you pass [4]int to this function, it will get a runtime err.
// However, go slices do not encode the size, meaning we can pass in any size to the func.
func Sum(arr [5]int) int {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	return sum
}

func SumSlice(slice []int) int {
	sum := 0

	for _, num := range slice {
		sum += num
	}

	return sum
}

// GO has varadic parameters, this is analgous to *args or js ...nums
// GO has named return values, they are trased as variables defined at the top of the function.
// A return statement without arguments, returns the named return values by default, known as a "naked return"
func SumAll(slices ...[]int) (sliceSums []int) {
	for _, slice := range slices {
		sum := SumSlice(slice)
		sliceSums = append(sliceSums, sum)

	}

	return sliceSums
}

func SumRecursive(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return nums[0] + SumRecursive(nums[1:])

}

// Slices can be 'sliced!'
// The syntax is slice[low:high], if you omit the value on one of the sides of :, it captures everything
// [1:] takes everying starting from the 1st index, [1:4] takes everything from the first, to the 3rd index (not including 4)
func SumAllTails(slices ...[]int) (sumOfTails []int) {
	for _, slice := range slices {
		if len(slice) == 0 {
			sumOfTails = append(sumOfTails, 0)
		} else {
			sum := SumSlice(slice[1:])
			sumOfTails = append(sumOfTails, sum)

		}
	}

	return sumOfTails
}
