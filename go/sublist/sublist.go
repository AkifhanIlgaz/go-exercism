package sublist

import (
	"reflect"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {

	if len(l1) < len(l2) {
		if isSubOrSuperlist(l1, Windows(l2, len(l1))) {
			return RelationSublist
		}
	} else if len(l1) > len(l2) {
		if isSubOrSuperlist(l2, Windows(l1, len(l2))) {
			return RelationSuperlist
		}
	} else {
		if reflect.DeepEqual(l1, l2) {
			return RelationEqual
		}
	}

	return RelationUnequal
}

func Windows(arr []int, window_size int) [][]int {
	w := [][]int{}

	for i := 0; i < len(arr)-window_size+1; i++ {
		w = append(w, arr[i:i+window_size])
	}

	return w
}

func isSubOrSuperlist(arr []int, windows [][]int) bool {
	for _, window := range windows {
		if reflect.DeepEqual(window, arr) {
			return true
		}
	}

	return false
}
