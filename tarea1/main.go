package main

import (
	"fmt"
)

func main() {
	data := readData("IRIS.csv")
	normalize(data)
	for i:=10; i< 50; i++{
		for j:= 5; j< 50; j++{
			_,cost := solve(data, 500,i,j)
			fmt.Println(cost,i,j)
		}
	}


}
