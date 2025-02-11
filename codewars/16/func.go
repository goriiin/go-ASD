package main

func Josephus(items []interface{}, k int) []interface{} {
	result := make([]interface{}, 0, len(items))
	index := 0

	for len(items) > 0 {
		index = (index + k - 1) % len(items)
		result = append(result, items[index])
		items = append(items[:index], items[index+1:]...)
	}

	return result
}
