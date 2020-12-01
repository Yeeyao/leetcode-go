package main

import "fmt"

// 需要减去公倍数部分
func main() {
	const maxNum = 1000
	const firstDivisor = 3
	const secondDivisor = 5

	var lessdivisor int

	sum := 0

	if firstDivisor > secondDivisor {
		lessdivisor = secondDivisor
	} else {
		lessdivisor = firstDivisor
	}

	maxtimes := maxNum / lessdivisor

	for i := 1; i < maxtimes+1; i++ {
		first := firstDivisor * i
		second := secondDivisor * i
		if first < maxNum {
			sum += first
		}
		if second < maxNum {
			sum += second
			// 这里减去公倍数
			if second%firstDivisor == 0 {
				sum -= second
			}
		}
	}

	fmt.Println(sum)
}
