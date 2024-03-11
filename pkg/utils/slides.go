package utils

// Map applies a function to each element of a slice and returns a new slice with the results
func Map[T, U any](f func(T) U, arr []T) []U {
	var result []U
	for _, v := range arr {
		result = append(result, f(v))
	}
	return result
}
