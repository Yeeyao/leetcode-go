package main

import "fmt"

func main() {
	const MAXNUM = 4000000
	sum := 0
	var i, j int
	for i, j = 1, 2; j < MAXNUM; i, j = j, i+j {
		if j&1 == 0 {
			sum += j
		}
	}
	fmt.Println(sum)
}
