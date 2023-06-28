package main

import "fmt"

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	// 8.6
	//fmt.Println("unsorted:")
	//for k, v := range barVal {
	//	fmt.Printf("Key: %v, Value: %v / ", k, v)
	//}
	//keys := make([]string, len(barVal))
	//i := 0
	//for k := range barVal {
	//	keys[i] = k
	//	i++
	//}
	////sort.Strings(keys)
	//sort.Slice(keys, func(i, j int) bool {
	//	return keys[i] > keys[j]
	//})
	//fmt.Println()
	//fmt.Println("sorted:")
	//for _, k := range keys {
	//	fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	//}
	// 8.7
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
}
