package main

import "fmt"

func main() {
	fmt.Println(t(1))
	return 0
}

/*
	这里形参和返回值的命名重复了，会报错
*/
func t(i int) (i int) {
	defer i++
	return 1
}
