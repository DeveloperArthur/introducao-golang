package main

import (
	"fmt" 
	"os"
	"strconv"
)

const FAHRENHEIT = "fahrenheit";
const MILHAS = "milhas";
const CELSIUS = "celsius";
const QUILOMETROS = "quilometros";

func main(){
	fechaProgramaSeNaoFoiPassadoParametro();

	unidadeOrigem := getUltimoParametroPassado();
	valoresOrigem := getListaDeParametrosNumericos();
	
	unidadeDestino := getUnidadeDestinoComBaseNa(unidadeOrigem);

	for i, v := range valoresOrigem {
		valorOrigem := validaSeNumeroEhValido(v, i);
		valorDestino := converteMedidas(unidadeOrigem, valorOrigem);
		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino)
	}
}

func fechaProgramaSeNaoFoiPassadoParametro(){
	if len(os.Args) < 3 {
		fmt.Println("vc precisa passar os parametros para funcionar")
		os.Exit(1);
	}
}

func getUltimoParametroPassado() string {
	return os.Args[len(os.Args)-1];
}

func getListaDeParametrosNumericos() []string {
	return os.Args[1 : len(os.Args)-1];
}

func getUnidadeDestinoComBaseNa(unidadeOrigem string) string {
	if unidadeOrigem == CELSIUS {
		return FAHRENHEIT;
	} else if unidadeOrigem == QUILOMETROS {
		return MILHAS;
	} else {
		fmt.Println("%s não é uma unidade conhecida!", unidadeOrigem)
		os.Exit(1);
		return "";
	}
}

func validaSeNumeroEhValido(v string, i int) float64 {
	valorOrigem, err := strconv.ParseFloat(v, 64)

	if err != nil {
		fmt.Println("O valor %s na posição %d não é um número válido!\n", v, i)
		os.Exit(1)
	}

	return valorOrigem;
}

func converteMedidas(unidadeOrigem string, valorOrigem float64) float64 {
	if unidadeOrigem == CELSIUS {
		return valorOrigem * 1.8 + 32
	} else {
		return valorOrigem / 1.60934
	}
}