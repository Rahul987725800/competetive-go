package main

import (
	"first/maths"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"math"
)

func main() {
	fmt.Println("ghanta van cho")
	fmt.Println(maths.Add(1, 4))
	fmt.Println(maths.Max(9, 10))
	fmt.Println(randomdata.SillyName())
	fmt.Println(math.MaxInt32)
}