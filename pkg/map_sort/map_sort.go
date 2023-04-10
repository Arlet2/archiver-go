package map_sort

import "sort"

type ordered interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | uintptr
}
type compare[V any] func(i, j V) bool

// SortMapByValues sorts map by values in some order (use isAscending for choosing). Sort is stable.
// Return sorted keys
func SortMapByValues[K comparable, V ordered](m map[K]V, isAscending bool) []K {
	var compare compare[V]
	if isAscending {
		compare = func(i, j V) bool {
			return i < j
		}
	} else {
		compare = func(i, j V) bool {
			return i > j
		}
	}

	return SortMapByValuesUsingCompare(m, compare)
}

// SortMapByKeys sorts map by keys in some order (use isAscending for choosing). Sort is stable.
// Return sorted keys
func SortMapByKeys[K ordered, V any](m map[K]V, isAscending bool) []K {
	var compare compare[K]
	if isAscending {
		compare = func(i, j K) bool {
			return i < j
		}
	} else {
		compare = func(i, j K) bool {
			return i > j
		}
	}

	return SortMapByKeysUsingCompare(m, compare)
}

// SortMapByKeysUsingCompare sorts map by keys with using client compare function.
// Compare function is compare[V any] func(i, j V) bool. Return sorted keys
func SortMapByKeysUsingCompare[K comparable, V any](m map[K]V, compare compare[K]) []K {
	keys := make([]K, 0)

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return compare(keys[i], keys[j])
	})

	return keys
}

// SortMapByValuesUsingCompare sorts map by values with using client compare function.
// Compare function is compare[V any] func(i, j V) bool. Return sorted keys
func SortMapByValuesUsingCompare[K comparable, V any](m map[K]V, compare compare[V]) []K {
	keys := make([]K, 0)

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return compare(m[keys[i]], m[keys[j]])
	})

	return keys
}
