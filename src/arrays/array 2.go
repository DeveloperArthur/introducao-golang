package main

import "fmt"

func main(){
	var multiA [2][2]int
	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, -2

	multiB := [2][2]int{{2, 13}, {-1, 6}}

	fmt.Println("Multi A: ", multiA)
	fmt.Println("Multi B: ", multiB)
}