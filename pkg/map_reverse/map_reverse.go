package map_reverse

// ReverseMap reverses keys and values in map
func ReverseMap[K comparable, V comparable](m map[K]V) (newMap map[V]K) {
	newMap = make(map[V]K, 0)

	for key, value := range m {
		newMap[value] = key
	}

	return
}
