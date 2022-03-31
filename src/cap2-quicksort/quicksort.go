package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	entrada := os.Args[1:]
	numeros := converteEntradaEmSliceDeNumeros(entrada);
	fmt.Println(quicksort(numeros))
}

func converteEntradaEmSliceDeNumeros(entrada []string) []int {
	numeros := make([]int, len(entrada))
	for i, n := range entrada {
	 	numero := validaSeNumeroEhValido(n);
		numeros[i] = numero
	}
	return numeros;
}

func validaSeNumeroEhValido(n string) int {
	numero, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("%s não é um número válido!\n", n)
		os.Exit(1)
	}
	return numero;
}

func quicksort(numeros []int) []int {
	comprimentoDosNumerosEhMaiorQue1 := len(numeros) <= 1
	if comprimentoDosNumerosEhMaiorQue1 { return numeros }

	n := criaCopiaDoSlice(numeros)

	indicePivo := len(n) / 2
	pivo := n[indicePivo]
	n = append(n[:indicePivo], n[indicePivo+1:]...)

	menores, maiores := particionar(n, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...);
}

func criaCopiaDoSlice(numeros []int) []int {
	n := make([]int, len(numeros))
	copy(n, numeros)
	return n;
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int){
	for _, n := range numeros{
		if n <= pivo{
			menores = append(menores, n)
		}else{
			maiores = append(maiores, n)
		}
	}

	return menores, maiores;
}