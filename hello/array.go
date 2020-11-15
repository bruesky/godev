package hello

import (
	"fmt"
)

func ArrayTest() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	var mySlice []int
	mySlice = make([]int, 8)
	mySlice = append(mySlice, 1)
	for index, value := range mySlice {
		fmt.Println(index, value)
	}

	var ranks map[string]int
	ranks = make(map[string]int)
	ranks["XG"] = 1
	ranks["RB"] = 2
	for key, value := range ranks {
		fmt.Println(key, value)
	}

}
