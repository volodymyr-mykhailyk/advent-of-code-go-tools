package science

func MinFromPairInt(v1, v2 int) int {
	if v1 <= v2 {
		return v1
	} else {
		return v2
	}
}

func MaxFromPairInt(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	} else {
		return v2
	}
}

func SortPairInt(v1, v2 int) (int, int) {
	if v1 <= v2 {
		return v1, v2
	} else {
		return v2, v1
	}
}
