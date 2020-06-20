package main

import (
	"fmt"
	"github.com/FranAmeri99/meli_ubp/tp3hello"
	"github.com/FranAmeri99/meli_ubp/tp4Calculdora"
	"github.com/FranAmeri99/meli_ubp/tp5ConsumoApis"
	"github.com/FranAmeri99/meli_ubp/tp6CalculadoraV2"
	"github.com/FranAmeri99/meli_ubp/tp7CalculadoraTest"
	"github.com/FranAmeri99/meli_ubp/tp8CalculadoraV3"
)

func main(){
	var switchN int64
	fmt.Printf("Ingrese el numero para correr el tp correspondiente:\n")
	fmt.Printf("3->Introduccion a Go\n")
	fmt.Printf("4->Variables y estructuras en Go\n")
	fmt.Printf("5->Consumo de APIs en Go\n")
	fmt.Printf("6->Concurrencia y Paralelirmos en Go\n")
	fmt.Printf("7->Testing\n")
	fmt.Printf("8->Creacion de una \n")
	fmt.Scan(&switchN)
	switch switchN {
	case 3:
		tp3hello.HelloWord()
		break
	case 4:
		tp4Calculdora.ldora.Calculadora()
		break
	case 5:
		tp5ConsumoApis.ConsumoApi()
		break
	case 6:
		tp6CalculadoraV2.CalculadoraHilos()
		var entrada string
		fmt.Scan(&entrada)
		break
	case 7:
		tp7CalculadoraTest.CalculadoraTest()
		break
	case 8:
		tp8CalculadoraV3.CalculadoraApi()
		break
	}
}
