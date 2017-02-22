package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const number = 338033

	startBasic := time.Now()
	fmt.Printf("basic Prime = %v, in %v\n", basic(number), time.Since(startBasic))

	startThreaded := time.Now()
	fmt.Printf("threaded Prime = %v, in %v\n", threaded(number), time.Since(startThreaded))

	startRecursive := time.Now()
	fmt.Printf("recursive Prime = %v, in %v\n", recursive(number, 2), time.Since(startRecursive))

	fmt.Scanln()
}

func basic(num int) (isPrime bool) {
	isPrime = true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	return
}

func recursive(num int, val int) (isPrime bool) {
	if val >= num {
		isPrime = true
		return
	}
	if num%val != 0 {
		isPrime = recursive(num, val+1)
	}
	return
}

func threaded(num int) (isPrime bool) {
	isPrime = true
	numbers := make(chan int)
	go generateNumber(numbers)
	answer := make(chan bool)
	go func() {
		for {
			created := <-numbers
			if created >= num {
				answer <- true
				break
			}
			if num%created == 0 {
				answer <- false
				break
			}
		}
	}()
	isPrime = <-answer
	return
}

func findAllThreaded(num int) (isPrime bool) {
	runtime.GOMAXPROCS(4)
	numbers := make(chan int)
	go generateNumber(numbers)

	for {
		number := <-numbers
		if number == num {
			isPrime = true
			break
		} else if number > num {
			isPrime = false
			break
		}
		fmt.Println(number)
		newFilters := make(chan int)
		go determinePrime(numbers, newFilters, number)
		numbers = newFilters
	}
	return
}

func determinePrime(in chan int, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func generateNumber(numbers chan int) {
	for i := 2; ; i++ {
		numbers <- i
	}
}
