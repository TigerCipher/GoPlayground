package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100
	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press <enter> to continue...")
	scanner.Scan()
	// goal, err := strconv.Atoi(scanner.Text())
	// if err == nil {
	
	for {
		guess := (low + high) / 2
		fmt.Println("Is the number", guess, "?")
		fmt.Println("Is that...")
		fmt.Println("a: too high?")
		fmt.Println("b: too low?")
		fmt.Println("c: correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		}else if response == "b"{
			low = guess + 1
		}else if response == "c" {
			fmt.Println("I am the victor!")
			break
		}else{
			fmt.Println("Invalid response")
		}
	}
	// }
}