package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("Hello")

	fmt.Println("The time is", time.Now())

	fmt.Println("My random number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(25))

	fmt.Printf("Pi is %g \n", math.Pi)
}