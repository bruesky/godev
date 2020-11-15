package hello

import "fmt"

func sum(channel chan int) {
	x := 0
	for i := 0; i < 10; i++ {
		x += i
	}
	channel <- x
}

func muti(channel chan int) {
	x := 1
	for i := 1; i < 10; i++ {
		x *= i
	}
	channel <- x
}

func ChanTest() {
	var myChan chan int
	myChan = make(chan int)
	var myChan2 chan int
	myChan2 = make(chan int)
	go sum(myChan)
	go muti(myChan2)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan2)
}
