package main

import (
	"fmt"

	"github.com/learn/algorithm/mySlice"
)

func main() {

	// s := []int8{1, 2, 3, 2, 3}
	// myMap := make(map[int]int)
	// for _,key := range s {

	// 	_,ok :=myMap[key]
	// 	if ok {
	// 		continue
	// 	}
	// 	myMap[i] = key

	// }

	//count := appearCount(s)

	//count := palindromenumber.APPEARCOUNT(s)
	// count := palindromenumber.APPEARCOUNT2(s)
	// fmt.Println(count)
	// strs := []string{"flower", "flow", "floight"}
	// result := str.STRMETHOD2(strs)
	// fmt.Println(result)

	// intArr := []uint64{1, 2, 3}

	// intArr = myInt.MYINTMETHOD(intArr)
	// fmt.Println(intArr)
	// arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// count := myInt.MYINTMETHOD2(arr)
	// fmt.Println(count)

	// interval := [][]int{{4, 7}, {2, 6}, {8, 10}, {15, 18}}
	// result := mySlice.MERGE(interval)
	// fmt.Println(result)
	intArr := []int{2, 7, 11, 15}
	target := 9
	a, b := mySlice.SUMFORTWONUM(intArr, target)
	fmt.Println(intArr[a])
	fmt.Println(intArr[b])
}
