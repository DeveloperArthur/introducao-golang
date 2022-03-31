package main

import "fmt"

type Carro struct {
	Name    string
	valores [3]int
}

func (c *Carro) andou() {
	//sem * nome é alterado apenas nesse escopo
	//com * ele vai mudar o nome q esta no meu objeto, pois estara apontando para o mesmo local na memoria
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func (c *Carro) alteraValor() {
	c.valores[0] = 500
}

func main() {
	carro := Carro{}
	carro.Name = "Punto"
	carro.valores = [3]int{100, 200, 300}

	carro.andou()
	carro.alteraValor()

	fmt.Println(carro.Name)
	fmt.Println(carro.valores)
}
