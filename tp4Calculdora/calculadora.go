package tp4_calculdora

import (
	"fmt"

	"github.com/FranAmeri99/meli_ubp/tp4Calculdora/operaciones"
)

func Calculadora()  {
	var num1 float64
	var num2 float64
	var op string
	var operandos []float64
	var opcion string
	var bug bool
	var offset int
	fmt.Printf("Ingrese la operación:\n")
	fmt.Printf("Sumar\n")
	fmt.Printf("Restar\n")
	fmt.Printf("Multiplicar\n")
	fmt.Printf("Dividir\n")
	fmt.Scan(&op)
	fmt.Printf("Ingrese 2 variables: \n")
	fmt.Scan(&num1, &num2)
	operandos = []float64{num1,num2}
	fmt.Printf("Desea buguear los resultados?: (Y/S)")
	fmt.Scan(&opcion)
	if(opcion == "S" || opcion == "Y" || opcion == "s" || opcion == "y"){
		bug = true
		fmt.Printf("Ingrese OFFSET:")
		fmt.Scan(&offset)
	}
	switch op {
	case "sumar":
		fmt.Print(operaciones.Sumar(bug,offset,operandos...))
	case "restar":
		fmt.Print(operaciones.Restar(bug,offset,operandos...))
	case "multiplicar":
		fmt.Print(operaciones.Multiplicar(bug,offset,operandos...))
	case "dividir":
		fmt.Print(operaciones.Division(bug,offset,operandos...))
	default:
		fmt.Print("Operracion inválida.")
	}
}
