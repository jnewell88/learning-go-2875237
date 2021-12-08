package main

import (
	"fmt"
	"sort"
)

func main() {

	//Map is unordered list of key value pairs (Hash)
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"
	fmt.Println(states)

	//iterate through map
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	//List in alphabetical order
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	//This list is sortable, so sort it to get alphabetical
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	//Print out result of states (with the ordered key indexes ordering the list)
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
