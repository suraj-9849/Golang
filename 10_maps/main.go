package main

import (
	"fmt"
	"maps"
	"slices"
)

// maps-> hashTables, Object, dict (Key Value Pair)
func main() {
	fmt.Println("Maps")
	m := make(map[string]string)
	// setting an element:
	m["name"] = "golang"
	m["area"] = "backend"
	fmt.Println(m["name"], m["area"])
	// if no key exists then we get the empty value:
	fmt.Println(m["phone"])

	// to delete the element from the Map:
	delete(m, "name")

	fmt.Println(m)
	// to clear the map :
	clear(m)
	fmt.Println(m)

	// other way of creating the Map:
	mp := map[string]int{"price": 2121, "age": 23, "phones": 232}
	fmt.Println(mp)

	// How to check whether the element is present in the Map or not:
	k, price := m["price"]
	if price {
		fmt.Println(k, price)
	} else {
		fmt.Println("Doesn't exists")
	}

	// There is the Package for the Map: maps
	mp1 := map[string]int{"price": 2345, "age": 20}
	mp2 := map[string]int{"price": 2345, "age": 20}
	fmt.Println(maps.Equal(mp1, mp2))
	fmt.Println(slices.Collect(maps.Keys(mp1)))
	fmt.Println(slices.Collect(maps.Keys(mp2)))
	fmt.Println(slices.Collect(maps.Values(mp1)))
	fmt.Println(slices.Collect(maps.Values(mp2)))
}
