// traduzir esse algoritmo em go: https://github.com/DeveloperArthur/algoritmos-guias-anotacoes-uteis/blob/main/listaUnificada.java
// colocar cliente em um arquivo separado
// alterar a lista com ponteiros

package main

import (
	"strings"
	"fmt"
)

type Cliente struct {
	Rg string
	Nome string
	ValorParaPagar float64
}

func main(){
	var clientes = []Cliente {
		{ "12345", "Sofia", 20.0 },
		{ "54321", "Sami", 40.0 },
		{ "12345", "Blade", 20.0 },
		{ "12345", "Sofia", 60.0 },
	}

	fmt.Println(clientes)
	unificaLista(clientes)
}

func mostraClientes(clientes []Cliente){
	for _,cliente := range clientes {
		fmt.Println(cliente.Rg, " ", cliente.Nome, " ", cliente.ValorParaPagar)
	}
}

func unificaLista(clientes []Cliente) {
	listaUnificada := make(map[string]Cliente)

	for _,cliente := range clientes {
		chave := strings.ToLower(cliente.Rg) + strings.ToLower(cliente.Nome)
		valorExistente, possuiChave := listaUnificada[chave]	
		if possuiChave {
			var valorNovo = valorExistente.ValorParaPagar + cliente.ValorParaPagar
			valorExistente.ValorParaPagar = valorNovo
			delete(listaUnificada, chave)
			listaUnificada[chave] = valorExistente
		} else {
			listaUnificada[chave] = cliente
		}
	}

	fmt.Println("------------------")
	fmt.Println(listaUnificada)
}