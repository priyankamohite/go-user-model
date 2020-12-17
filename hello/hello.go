package main

import (
	"fmt"
	//"math"
	//"time"
)

type Vertex struct {
	x, y int
}

func main() {
	//fmt.Println("Hello, its working")
	//fmt.Println("current time is : ", time.Now())
	//fmt.Printf("sqrt--------- %g \n", math.Sqrt(4))

	//fmt.Println(getAddition(22, 33))

	//fmt.Println(swapNumbers(22, 33))
	// fmt.Println(getSplit())
	//printNumbers(0, 10)

	v := Vertex{22, 33}
	fmt.Println(v.checkMethods())

}

func getAddition(x, y int) int {

	return x + y
}

func swapNumbers(x, y int) (int, int) {
	return y, x
}

func printNumbers(low, upper int) {

	for number := low; number < upper+1; number++ {
		fmt.Println(number, upper)
	}
}

func (V Vertex) checkMethods() int {

	return V.x - V.y

}
