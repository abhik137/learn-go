package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x // send value into c
		x, y = y, x+y
	}
	close(c) // required to terminate range over channel
}

func main() {
	c := make(chan int, 10)
	/* The cap built-in function returns the capacity of v, according to its type:
	 * Channel: the channel buffer capacity, in units of elements; */
	go fibonacci(cap(c), c)
	// range over channel
	for i := range c {
		fmt.Println(i)
	}
}
