package science

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func SumElementsInt(elements []int) int {
	total := 0
	for _, value := range elements {
		total += value
	}
	return total
}

func MultiplyElementsInt(elements []int) int {
	total := 1
	for _, value := range elements {
		total *= value
	}
	return total
}
