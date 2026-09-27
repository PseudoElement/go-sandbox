package main

func copyMap[K comparable, V comparable](m map[K]V) map[K]V {
	newMap := make(map[K]V, len(m))
	for key, value := range m {
		newMap[key] = value
	}
	return newMap
}
