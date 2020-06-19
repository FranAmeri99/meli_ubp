package operaciones

type estructura struct {
	Bug bool
	Offset int
	Operandos []float64
	Operacion string
	Resultado float64
}
func Division(bug bool,offset int,operandos ...float64) (estructura) {
	resultado := operandos[0] / operandos[1]
	if(bug){
		resultado += float64(offset)
	}
	return  estructura{
		Bug: bug,
		Offset: offset,
		Operandos: operandos,
		Operacion: "Division",
		Resultado: resultado,
	}
}
func Restar(bug bool,offset int,operandos ...float64) (estructura) {
	resultado := operandos[0] - operandos[1]
	if(bug){
		resultado += float64(offset)
	}
	return  estructura{
		Bug: bug,
		Resultado: resultado,
		Offset: offset,
		Operacion: "resta",
		Operandos: operandos,
	}
}

func Sumar(bug bool, offset int, operandos ...float64) (estructura) {
	resultado :=operandos[0] + operandos[1]
	if(bug){
		resultado += float64(offset)
	}
	return  estructura{
		Bug: bug,
		Offset: offset,
		Operandos: operandos,
		Operacion: "suma",
		Resultado: resultado,
	}
}


func Multiplicar(bug bool,offset int,operandos ...float64) (estructura) {
	resultado := operandos[0] * operandos[1]
	if(bug){
		resultado += float64(offset)
	}
	return  estructura{
		Bug: bug,
		Offset: offset,
		Operandos: operandos,
		Operacion: "multiplicacion",
		Resultado: resultado,
	}
}

