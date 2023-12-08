package science

// GreatCommonDivisor find greatest common divisor via Euclidean algorithm
func GreatCommonDivisor[T Number](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LeastCommonMultiple find The Least Common Multiple (LCM) via GreatCommonDivisor
func LeastCommonMultiple[T Number](numbers []T) T {
	if len(numbers) > 1 {
		a := numbers[0]
		b := numbers[1]
		result := a * b / GreatCommonDivisor(a, b)

		return LeastCommonMultiple(append(numbers[2:], result))
	} else {
		return numbers[0]
	}
}
