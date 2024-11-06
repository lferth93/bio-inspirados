package main

import (
	"fmt"
)

func main() {
	data := readData("IRIS.csv")
	normalize(data)
	sol := solve(data, 1000)

	fmt.Println(sol)

}
