package util

type number interface {
	int | int8 | int16 | int32 | int64
}

func Abs[T number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
