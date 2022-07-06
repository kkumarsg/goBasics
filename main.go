package main

import (
	"fmt"
	"math/rand"
)

func multiply(x float64, y float64) (float64, float64) {
	return x * y, y
}

func main() {
	fmt.Println("Say Hello ")
	fmt.Println("Random number is ", rand.Int63n(100))
}
