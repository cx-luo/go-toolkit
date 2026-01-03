// Package maputil provides map manipulation utilities
package maputil

// Keys returns all keys from a map
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values from a map
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// ContainsKey checks if a map contains a key
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, exists := m[key]
	return exists
}

// GetOrDefault returns the value for a key, or a default value if the key doesn't exist
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, exists := m[key]; exists {
		return v
	}
	return defaultValue
}

// Merge merges multiple maps into one (later maps override earlier ones)
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// Filter filters a map based on a predicate function
func Filter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// Map applies a function to each key-value pair and returns a new map
func Map[K comparable, V any, R any](m map[K]V, mapper func(K, V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = mapper(k, v)
	}
	return result
}

// Invert inverts a map (swaps keys and values)
func Invert[K, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// IsEmpty checks if a map is empty
func IsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}

// IsNotEmpty checks if a map is not empty
func IsNotEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) > 0
}

// Clear removes all entries from a map
func Clear[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

// Copy creates a shallow copy of a map
func Copy[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

