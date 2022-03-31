package main

import "fmt"

func main(){
	numeros := [6]int{1,2,3,4,5,6}

	//fatiando e pegando uma nova colecao só que nao quero a primeira posicao
	slice := numeros[1:6]

	//enderecos de memoria sao iguais pq nao criei outro array com os mesmos numeros, criei um slice que aponta para o mesmo endereco de memoria, nisso ele tem os mesmos valores
	fmt.Println(&numeros[1])
	fmt.Println(&slice[0])

	fmt.Println(numeros) 
	fmt.Println(slice) 
	//a mudança afeta o slice pq ele tava apontando para o array
	numeros[1] = 0 
	fmt.Println(numeros) 
	fmt.Println(slice) 

	fmt.Println(numeros)
	recebeSlice(slice)
	fmt.Println(numeros)
}

//sintaxe para receber um slice
func recebeSlice(valores []int){
	valores[0] = 69
}