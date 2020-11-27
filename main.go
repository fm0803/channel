package main

import (
	"fmt"
	 "time"
)

func main() {
	ch := make(chan int, 100)

	go func() {
		for i:=0;i<10;i++ {
			ch <- i
		}
	}()

	go func ()  {
		for i:=0;i<10;i++ {
			a,ok := <- ch
			if !ok {
				fmt.Println("ch close")
				return
			}
			fmt.Println("a:",a)
		}
	}()

	close(ch)
	fmt.Println("OK")
	time.Sleep(time.Second * 3)
}
