package selectionSort

func Execute(list []int) []int {
	var resultList []int
	var smallestIndex int
	listLength := len(list)

	for i := 0; i < listLength; i++ {
		smallestIndex = findSmallest(list)
		resultList = append(resultList, list[smallestIndex])
		list = append(list[:smallestIndex], list[smallestIndex+1:]...)
	}

	return resultList
}

func findSmallest(list []int) int {
	tmpElement := list[0]
	smallestIndex := 0
	for i, element := range list {
		if tmpElement > element {
			tmpElement = element
			smallestIndex = i
		}
	}

	return smallestIndex
}
