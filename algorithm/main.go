package main

import (
	"fmt"

	"github.com/learn/algorithm/palindromenumber"
)

func main() {

	s := []int8{1, 2, 3, 2, 3}
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
	count := palindromenumber.APPEARCOUNT2(s)

	fmt.Println(count)
}
