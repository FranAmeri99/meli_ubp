package tp6_calculadora_hilos

import (
	"fmt"
	"time"
)

type estructura struct {
	Bug bool
	Offset int
	Operandos []float64
	Operacion string
	Resultado float64
}

var c1 chan estructura =  make(chan estructura)
var c2 chan estructura = make(chan estructura)
var c3 chan estructura =  make(chan estructura)
var c4 chan estructura =  make(chan estructura)

func CalculadoraHilos() {

	var num1, num2 float64
	var operandos []float64
	var opcion string
	var bug bool
	var offset int


	fmt.Printf("Ingrese dos numeros: \n")
	fmt.Scan(&num1, &num2)
	operandos = []float64{num1, num2}
	fmt.Printf("Desea buguear el resultado?: (presione S/Y)")
	fmt.Scan(&opcion)
	if (opcion == "S" || opcion == "s" || opcion == "Y" || opcion == "y") {
		bug = true
		fmt.Printf("Ingrese OFFSET del Bug:")
		fmt.Scan(&offset)
	}

	go func() {

		resultado := operandos[0] + operandos[1]
		if (bug) {
			resultado += float64(offset)
		}
		c1 <- estructura{
			Bug:       bug,
			Offset:    offset,
			Operandos: operandos,
			Operacion: "suma",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		resultado := operandos[0] - operandos[1]
		if (bug) {
			resultado += float64(offset)
		}
		c2 <- estructura{
			Bug:       bug,
			Offset:    offset,
			Operandos: operandos,
			Operacion: "resta",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)
	}()

	go func() {

		resultado := operandos[0] * operandos[1]
		if (bug) {
			resultado += float64(offset)
		}
		c3 <- estructura{
			Bug:       bug,
			Offset:    offset,
			Operandos: operandos,
			Operacion: "multiplicacion",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)
	}()

	go func() {

		resultado := operandos[0] / operandos[1]
		if (bug) {
			resultado += float64(offset)
		}
		c4<- estructura{
			Bug:       bug,
			Offset:    offset,
			Operandos: operandos,
			Operacion: "Division",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)
	}()

	go func() {
		for {
			select {
			case suma := <-c1:
				fmt.Println(suma)
			case resta := <-c2:
				fmt.Println(resta)
			case multiplicar := <-c3:
				fmt.Println(multiplicar)
			case dividir := <-c4:
				fmt.Println(dividir)
			}
		}
	}()
}
