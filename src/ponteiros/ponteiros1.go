package main

import "fmt"

//https://www.youtube.com/watch?v=-FiBp1OeZF0
func main(){
	//memoria do PC-tem endereço 0xc00001c088 (guarda valor 10) <--- variavel a <---- valor 10

	a := 10

	//mostra o endereço na memoria
	fmt.Println(&a)

	//ponteiro tem o valor do endereço de memória da variavel a, ou seja = 0xc00001c088
	var ponteiro *int = &a

	//mostra o endereço na memoria
	fmt.Println(ponteiro)

	//mostra o valor da variavel a
	fmt.Println(*ponteiro)

	//o valor 50 é armazenado no endereço de memória do ponteiro
	*ponteiro = 50

	//variavel a mudou pois armazenamos 50 no ponteiro e o ponteiro aponta para o endereço de memória de a
	fmt.Println(a)
}