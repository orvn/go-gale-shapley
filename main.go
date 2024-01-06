package main

import "fmt"

func main() {
	// Define preferences here (can be dynamic or read from external sources)
	groupAPref := [][]int{{3, 1, 2, 0}, {1, 0, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}}
	groupBPref := [][]int{{0, 1, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}, {0, 1, 2, 3}}

	partners, err := stableMatching(groupAPref, groupBPref)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display the results
	fmt.Println("Group B ", " Group A")
	for i, partner := range partners {
		fmt.Println(i, "\t", partner)
	}
}
