package main

import "fmt"

func main(){
	numeros := [6]int{1,2,3,4,5,6}

	//fatiando e pegando uma nova colecao só que nao quero a primeira posicao
	slice := numeros[1:6]

	//enderecos de memoria sao iguais slice, pq nao criei outro array com os mesmos numeros, 
	//criei um slice que aponta para o mesmo endereco de memoria, nisso ele tem os mesmos valores
	fmt.Println(&numeros[1]) //endereco de memoria 0xc000122008
	fmt.Println(&slice[0]) //endereco de memoria 0xc000122008

	fmt.Println(numeros) // [1 2 3 4 5 6]
	fmt.Println(slice) // [2 3 4 5 6]

	//se mudar o array, a mudança afeta o slice 
	//pq ele tava apontando para o array
	numeros[1] = 0 
	fmt.Println(numeros) //[1 0 3 4 5 6]
	fmt.Println(slice) //[0 3 4 5 6]

	//e o mesmo serve para o slice
	//se mudar o slice, afeta o array
	fmt.Println(numeros)
	recebeSlice(slice)
	fmt.Println(numeros)
	
	//outros exemplos: https://github.com/DeveloperArthur/golang-first-api-rest/blob/main/service/cliente_service.go
	//RUST: https://github.com/DeveloperArthur/introducao-rust/ownership/slice.rs
}

//sintaxe para receber um slice
func recebeSlice(valores []int){
	valores[0] = 69
}
