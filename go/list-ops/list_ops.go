package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	accumulator := initial

	for _, element := range s {
		accumulator = fn(accumulator, element)
	}

	return accumulator
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accumulator := initial

	for i := s.Length() - 1; i >= 0; i-- {
		accumulator = fn(s[i], accumulator)
	}

	return accumulator
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filteredList := IntList{}
	for _, element := range s {
		if !fn(element) {
			continue
		}

		filteredList = filteredList.Append([]int{element})
	}

	return filteredList
}

func (s IntList) Length() int {
	length := 0

	for range s {
		length++
	}

	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	for i, element := range s {
		s[i] = fn(element)
	}

	return s
}

func (s IntList) Reverse() IntList {
	sLength := s.Length()
	reversedList := make(IntList, sLength)

	normalIndex := 0
	for i := sLength - 1; i >= 0; i-- {
		reversedList[i] = s[normalIndex]
		normalIndex++
	}

	return reversedList
}

func (s IntList) Append(lst IntList) IntList {
	sLength := s.Length()
	lstLength := lst.Length()
	totalLength := sLength + lstLength

	newList := make(IntList, totalLength)

	for i, element := range s {
		newList[i] = element
	}

	for i, element := range lst {
		newList[sLength+i] = element
	}

	return newList
}

func (s IntList) Concat(lists []IntList) IntList {
	allLists := s

	for _, list := range lists {
		allLists = allLists.Append(list)
	}

	return allLists
}
