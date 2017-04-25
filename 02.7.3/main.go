package main

import "fmt"

func main() {
	//修改2为1就报错，修改2为3可以正常运行
	//todo why修改2为1就报错?
	c := make(chan int, 1)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
