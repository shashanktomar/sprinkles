package utils

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Limit(min int, max int, num int) int {
	if num < min {
		num = min
	}
	if num > max {
		num = max
	}
	return num
}
