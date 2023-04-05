package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var newInts Ints

	for _, num := range i {
		if filter(num) {
			newInts = append(newInts, num)
		}
	}

	return newInts
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var newInts Ints

	for _, num := range i {
		if !filter(num) {
			newInts = append(newInts, num)
		}
	}

	return newInts
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var newLists Lists

	for _, num := range l {
		if filter(num) {
			newLists = append(newLists, num)
		}
	}

	return newLists
}

func (s Strings) Keep(filter func(string) bool) Strings {

	var newStrings Strings

	for _, num := range s {
		if filter(num) {
			newStrings = append(newStrings, num)
		}
	}

	return newStrings
}
