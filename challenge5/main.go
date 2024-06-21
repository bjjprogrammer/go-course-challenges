package main

import "fmt"

func sum(numbers []int, ch chan int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	ch <- total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	middle := len(numbers) / 2
	firstHalf := numbers[:middle]
	secondHalf := numbers[middle:]

	ch := make(chan int)

	go sum(firstHalf, ch)

	go sum(secondHalf, ch)

	sum1 := <-ch
	sum2 := <-ch
	totalSum := sum1 + sum2

	fmt.Printf("La suma total es: %d\n", totalSum)
}
