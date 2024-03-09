package main

import (
	"fmt"
	"log"

)

func main() {
	//Long
	var names [2]string
	names[0] = "Shinnamaon"
	names[1] = "15001"
	fmt.Println(names)
	//Short
	labels := [2]string{"x", "y"}
	fmt.Println(labels)

	printTest(40, 20, subtract)

	num, err := simerr()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)

	// x := subtract(10, 5)
	// swapint(10, 50)
	// fmt.Println(x)
}

func printTest(x, y int, f func(int, int) int) {
	z := f(x, y)

	fmt.Println(z)
}

func subtract(x, y int) int {
	return x - y
}

func simerr() (int, error) {
	return 20, fmt.Errorf("error simulate %d", 1)
}
