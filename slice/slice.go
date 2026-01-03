// Package slice provides slice manipulation utilities
package slice

// Contains checks if a slice contains a value
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of value in slice, or -1 if not found
func IndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// Remove removes the first occurrence of value from slice
func Remove[T comparable](slice []T, value T) []T {
	index := IndexOf(slice, value)
	if index == -1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// RemoveAll removes all occurrences of value from slice
func RemoveAll[T comparable](slice []T, value T) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// Unique returns a new slice with unique values
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Filter filters a slice based on a predicate function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map applies a function to each element of a slice and returns a new slice
func Map[T any, R any](slice []T, mapper func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

// Reduce reduces a slice to a single value using a reducer function
func Reduce[T any, R any](slice []T, initial R, reducer func(R, T) R) R {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}

// Reverse reverses a slice
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}

// Chunk splits a slice into chunks of specified size
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return [][]T{slice}
	}
	chunks := make([][]T, 0, (len(slice)+size-1)/size)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// Flatten flattens a 2D slice into a 1D slice
func Flatten[T any](slices [][]T) []T {
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}
	result := make([]T, 0, totalLen)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// Intersect returns the intersection of two slices
func Intersect[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	for _, v := range slice2 {
		set[v] = true
	}
	result := make([]T, 0)
	for _, v := range slice1 {
		if set[v] {
			result = append(result, v)
		}
	}
	return Unique(result)
}

// Union returns the union of two slices
func Union[T comparable](slice1, slice2 []T) []T {
	return Unique(append(slice1, slice2...))
}

// Difference returns the difference of two slices (elements in slice1 but not in slice2)
func Difference[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	for _, v := range slice2 {
		set[v] = true
	}
	result := make([]T, 0)
	for _, v := range slice1 {
		if !set[v] {
			result = append(result, v)
		}
	}
	return result
}

// IsEmpty checks if a slice is empty
func IsEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// IsNotEmpty checks if a slice is not empty
func IsNotEmpty[T any](slice []T) bool {
	return len(slice) > 0
}

// First returns the first element of a slice, or zero value if empty
func First[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[0]
}

// Last returns the last element of a slice, or zero value if empty
func Last[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[len(slice)-1]
}

// Take returns the first n elements of a slice
func Take[T any](slice []T, n int) []T {
	if n <= 0 {
		return []T{}
	}
	if n >= len(slice) {
		return slice
	}
	return slice[:n]
}

// Skip returns a slice with the first n elements removed
func Skip[T any](slice []T, n int) []T {
	if n <= 0 {
		return slice
	}
	if n >= len(slice) {
		return []T{}
	}
	return slice[n:]
}

