package science

type Number interface {
	~int | ~int64
}

func ExecuteOperation[T Number](left T, operation string, right T) T {
	switch operation {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		return 0
	}
}
