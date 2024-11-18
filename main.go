package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	x := 5
	y := 5
	sum := x + y

	a := []int{0,1,2,3,4}
	a = append(a, 500)

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	fmt.Println(vertices["triangle"])

	fmt.Println("Hello, the sum is ", sum)

	fmt.Println("Here is an array of numbers", a)

	fmt.Println("The time is", time.Now())

	fmt.Println("My random number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(25))

	fmt.Printf("Pi is %g \n", math.Pi)

	fmt.Println("Hello JL")

	fmt.Scanln()
}