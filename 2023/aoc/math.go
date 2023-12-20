package aoc

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
