package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

//
// var map map[keyType]valueType
//
func main() {

}

func sortMap() {
	fmt.Println("before:")
	fmt.Println(barVal)
	keys := make([]string, len(barVal))
	var i = 0
	for key, _ := range barVal {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	fmt.Println("after:")
	for _, k := range keys {
		fmt.Println(k, ":", barVal[k])
	}
}

func ttttttt() {
	var m1 = make(map[string]int)
	m1["hello"] = 1
	m1["world"] = 2

	val, isPresent := m1["hello"]
	fmt.Println(val)
	fmt.Println(isPresent)
	//fmt.Println(ok)
	if val, isPresent := m1["hello"]; isPresent {
		fmt.Println(val)
	}
}
