package hello

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func Test() {
	fmt.Println(reflect.TypeOf(43))
	fmt.Println(reflect.TypeOf(43.00))
	fmt.Println(reflect.TypeOf(43 == 43))
	fmt.Println(reflect.TypeOf("43"))

	var quantity, num int = 100, 99
	num = 98
	fmt.Println(quantity)
	fmt.Println(num)
	index := 0
	fmt.Printf("%v %T\n", index, index)

	fmt.Println(time.Now().Year())

	broken := "G# R#cks"
	rep := strings.NewReplacer("#", "o")
	fmt.Println(rep.Replace(broken))

	const ONE = 1
	fmt.Println(ONE)

}
