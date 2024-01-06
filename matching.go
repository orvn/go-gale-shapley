package main

import "fmt"

// Getter of index of element in preference list
func getPreferenceIndex(preferences [][]int, a, b int) int {
	for i, pref := range preferences[a] {
		if pref == b {
			return i
		}
	}
	return -1
}

// Gale-Shapley algorithm for stable matching between two groups
func stableMatching(groupAPref, groupBPref [][]int) ([]int, error) {
	if len(groupAPref) != len(groupBPref) {
		return nil, fmt.Errorf("group sizes do not match")
	}

	n := len(groupAPref)
	bPartner := make([]int, n)
	for i := range bPartner {
		bPartner[i] = -1
	}

	aFree := make([]bool, n)

	freeCount := n

	for freeCount > 0 {
		var a int
		for a = 0; a < n; a++ {
			if !aFree[a] {
				break
			}
		}

		for i := 0; i < n && !aFree[a]; i++ {
			b := groupAPref[a][i]

			if bPartner[b] == -1 {
				bPartner[b] = a
				aFree[a] = true
				freeCount--
			} else {
				a1 := bPartner[b]
				if getPreferenceIndex(groupBPref, b, a) < getPreferenceIndex(groupBPref, b, a1) {
					bPartner[b] = a
					aFree[a] = true
					aFree[a1] = false
				}
			}
		}
	}

	return bPartner, nil
}
