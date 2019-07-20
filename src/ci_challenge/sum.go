package main

import (
	"fmt"
	"math"
)

func soma( ) string {

	inicio := 0.001

	for j := 0; j <= 1000000; j++ {

		inicio = inicio + math.Sqrt(inicio)
	}
	
	return "Code.education Rocks!"
}

func main() {

	fmt.Println(soma())
}