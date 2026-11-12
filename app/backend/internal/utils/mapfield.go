package utils

func MapField[T, R any](items []T, fn func(T) (R, bool)) []R {
	out := make([]R, 0, len(items))
	for _, item := range items {
		if v, ok := fn(item); ok {
			out = append(out, v)
		}
	}
	return out
}

// MapValues mengubah/memfilter semua value dari map menjadi slice baru.
// fn mengembalikan (hasil, ok) — jika ok == false, item di-skip.
func MapValues[K comparable, V, R any](m map[K]V, fn func(K, V) (R, bool)) []R {
	out := make([]R, 0, len(m))
	for k, v := range m {
		if r, ok := fn(k, v); ok {
			out = append(out, r)
		}
	}
	return out
}

// FilterMap mengembalikan map baru berisi entry yang lolos kondisi fn.
func FilterMap[K comparable, V any](m map[K]V, fn func(K, V) bool) map[K]V {
	out := make(map[K]V)
	for k, v := range m {
		if fn(k, v) {
			out[k] = v
		}
	}
	return out
}

// FindValue mencari satu value pertama yang memenuhi kondisi fn.
// Karena urutan iterasi map di Go tidak terjamin, "pertama" di sini
// artinya "yang pertama ditemukan saat iterasi", bukan berurutan.
func FindValue[K comparable, V any](m map[K]V, fn func(K, V) bool) (V, bool) {
	for k, v := range m {
		if fn(k, v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// AnyMatch mengecek apakah ada minimal satu entry yang memenuhi kondisi fn.
func AnyMatch[K comparable, V any](m map[K]V, fn func(K, V) bool) bool {
	for k, v := range m {
		if fn(k, v) {
			return true
		}
	}
	return false
}

// CountMatch menghitung jumlah entry yang memenuhi kondisi fn.
func CountMatch[K comparable, V any](m map[K]V, fn func(K, V) bool) int {
	n := 0
	for k, v := range m {
		if fn(k, v) {
			n++
		}
	}
	return n
}

func Unique(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	out := make([]string, 0, len(in))
	for _, s := range in {
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			out = append(out, s)
		}
	}
	return out
}
