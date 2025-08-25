package palindromenumber

func APPEARCOUNT(array []int8) int8 {
	var count int8 = 0
	for i := 0; i < len(array); i++ {
		for _, value := range array {

			if value == array[i] {
				count++
			}
		}
		if count == 1 {
			return int8(array[i])
		}
	}
	return 0
}
func APPEARCOUNT2(array []int8) int8 {
	myMap := make(map[int8]int8)
	for i := 0; i < len(array); i++ {

		myMap[array[i]]++

	}
	for key, value := range myMap {
		if value == 1 {
			return key
		}
	}
	return 0
}
