package map_sort

import "sort"

type ordered interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	float32 | float64 |
	string | uintptr
}

// SortMapByValues sort map by values in some order (use isAscending for choosing). Sort is stable.
// Return sorted keys
func SortMapByValues[K comparable, V ordered](m map[K]V, isAscending bool) []K {
	keys := make([]K, 0)

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if isAscending {
			return m[keys[i]] < m[keys[j]]
		} else {
			return m[keys[i]] > m[keys[j]]
		}

	})

	return keys
}

// SortMapByKeys sort map by keys in some order (use isAscending for choosing). Sort is stable.
// Return sorted keys
func SortMapByKeys[K ordered, V any](m map[K]V, isAscending bool) []K {
	keys := make([]K, 0)

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if isAscending {
			return keys[i] < keys[j]
		} else {
			return keys[i] > keys[j]
		}

	})

	return keys
}
