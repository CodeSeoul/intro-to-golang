package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go sayHi()
	}
	sayHi()
	fmt.Println("Finished?!")

	time.Sleep(1 * time.Second)
	fmt.Println("Finished after waiting")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("x:", x, "-- y:", y, "-- x+y:", x+y)

	// Let's limit the channel to 2 values
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// Uncomment the below and see what happens
	// ch <- 3
	fmt.Println("channel:", <-ch)
	fmt.Println("channel:", <-ch)

	// What if you want to continuously read from a
	// channel until it's finished sending?
	c = make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Done!")

	// Selecting channels is good for graceful shutdown
	c = make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciWithQuit(c, quit)
}

func sayHi() {
	randomWait := rand.IntN(1000)
	time.Sleep(time.Duration(randomWait))
	fmt.Println("Hi after", strconv.Itoa(randomWait)+"!")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciWithQuit(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
