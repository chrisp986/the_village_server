package main

import (
	"fmt"
	"math/rand"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func calcResources() {
	fmt.Println("Calculating resources...")
}
