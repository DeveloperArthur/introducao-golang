package main

import "fmt"

func main(){
	var a [3]int
	numeros := [5]int{1,2,3,4,5}
	primos := [...]int{2,3,5,7,11,13}
	nomes := [2]string{"Arthur", "Gabu"}

	teste := [...]int

	fmt.Println(a, numeros, primos, nomes, teste)
}