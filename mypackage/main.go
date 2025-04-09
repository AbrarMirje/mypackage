package mypackage

// sum add unlimited number of values of type int
func Sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
