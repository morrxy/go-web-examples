package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

//todo 为什么结果是-5 17 12 而不是17 -5 12？
func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go func() {
		t1 := 5000
		fmt.Println("first half", t1)
		time.Sleep(time.Duration(t1) * time.Microsecond)
		sum(a[:len(a)/2], c)
	}()

	go func() {
		t2 := 1
		fmt.Println("second half", t2)
		time.Sleep(time.Duration(t2) * time.Microsecond)
		sum(a[len(a)/2:], c)
	}()

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
