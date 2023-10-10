package main

import "fmt"

func main() {

	// var a int = 5
	// var a = 5 // implicit int
	a := 5 // short cut for the above

	// var f1 float32
	// var f2 float64

	pi := 3.14159

	fmt.Println(a + int(pi))
	fmt.Println(float64(a) + pi)

	var b uint8 = 250
	var c uint8 = 10
	var res uint8
	res = b + c
	fmt.Println(res) // Prints 4 because of overflow. 250 + 10 = 260 but max is 255, it wraps around to 0. So what should be 256 is 0, etc

	f1 := 2.0
	f2 := 3.0
	fres := f1 / f2

	fmt.Printf("%.20f\n", fres) // print 20 decimal places

	// var flag bool = true
	flag := true
	fmt.Println(flag)

	helloworld := "Hello World!"
	fmt.Println(helloworld)

	// infinite loop - think while(true) in C++
	/*for {

	} */

	// Equivalent of: while(counter < 10)
	counter := 0
	for counter < 10 {
		fmt.Println("Counter: ", counter)
		counter++
	}
	// Equivalent of: for(int i = 0; i < 10; ++i)
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d, ", i)
	}
	fmt.Println()

	flag = false
	if flag {
		fmt.Println("Flag was true")
	} else {
		fmt.Println("Flag was false")
	}

	if !flag {
		fmt.Println("Goodbye world")
	} else {
		fmt.Println("Oh, hello world")
	}

	x := 3
	if x > 5 {
		fmt.Println("X is greater than 5")
	} else if x < 5{
		fmt.Println("X is less than 5")
	} else {
		fmt.Println("X is 5")
	}
}