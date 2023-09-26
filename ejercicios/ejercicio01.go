package ejercicios

import (
	"fmt"
	"strconv"
)

func ConvierteTextoNumero(numero string) (int, error) {
	resultado, err := strconv.Atoi(numero)

	if resultado > 100 {
		fmt.Println("Es mayor a 100")
	} else {
		fmt.Println("Es menor a 100")
	}

	return resultado, err
}
