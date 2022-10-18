package main

// It will be replaced by constraints
type Signed interface {
	int | int8 | int16 | int32 | int64
}

// Generics
func GMin[T Signed](list []T) T {
	if len(list) == 0 {
		var empty T
		return empty
	}

	m := list[0]

	for _, v := range list {
		if v < m {
			m = v
		}
	}
	return m
}

func Min(list []int) int {
	if len(list) == 0 {
		return 0
	}
	m := list[0]
	for _, l := range list {
		if l < m {
			m = l
		}
	}

	return m
}

func GMax[T Signed](list []T) T {
	if len(list) == 0 {
		var empty T
		return empty
	}

	m := list[0]

	for _, v := range list {
		if v > m {
			m = v
		}
	}
	return m

}
func Max(list []int) int {
	if len(list) == 0 {
		return 0
	}

	m := list[0]
	for _, l := range list {
		if l > m {
			m = l
		}
	}

	return m
}
