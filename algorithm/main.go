package main

import (
	"fmt"

	"github.com/learn/algorithm/str"
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
	strs := []string{"flower", "flow", "floight"}
	result := str.STRMETHOD2(strs)
	fmt.Println(result)

}
