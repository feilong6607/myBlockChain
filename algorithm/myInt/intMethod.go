package myInt

func MYINTMETHOD(intArray []uint64) []uint64 {

	for i := 0; i < len(intArray); i++ {

		intArray[i] = intArray[i] + 1
	}
	return intArray

}
func MYINTMETHOD2(intArray []int) int {

	var k1 int = 0

	var k2 int = 1

	for k2 < len(intArray) {

		if intArray[k1] == intArray[k2] {

			k2++

		} else {
			k1++
			intArray[k1] = intArray[k2]
		}
	}
	return k1 + 1

}
