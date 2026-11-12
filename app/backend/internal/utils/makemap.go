package utils

func MapSlice[T ~string | ~int | ~int64 | ~float64](keys ...T) []T {
	var res []T
	for _, key := range keys {
		res = append(res, key)
	}
	return res
}
