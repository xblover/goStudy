package main

import "fmt"

func main() {
	var i,j int = 1,2
	//在函数中，简洁赋值语句 := 可在类型明确的地方代替var声明
	k := 3
	c,python, java := true, false, "no!"

	fmt.Println(i,j,k,c,python,java)

}