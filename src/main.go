package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		//Declaración de constantes
		const pi float64 = 3.14
		const pi2 = 3.14

		fmt.Println("pi:", pi)
		fmt.Println("pi2:", pi)

		//Declaración de variables enteras
		base := 12
		var altura int = 14
		var area int
		fmt.Println(base, altura, area)

		//Zero Values
		var a int
		var b float64
		var c string
		var d bool

		fmt.Println(a, b, c, d)

		//Area del Cuadrado
		const baseCuadrado = 10
		areaCuadrado := baseCuadrado * baseCuadrado
		fmt.Println("Area Cuadrado:", areaCuadrado)

		x := 10
		y := 50

		//Suma
		result := x + y
		fmt.Println("Suma:", result)

		//Resta
		result = y - x
		fmt.Println("Resta:", result)

		//Multiplicacion
		result = x * y
		fmt.Println("Multiplicacion:", result)

		//Division
		result = y / x
		fmt.Println("Division:", result)

		//Modulo
		result = y % x
		fmt.Println("Modulo:", result)

		//Incremental
		x++
		fmt.Println("Incremental:", x)

		//Decremental
		x--
		fmt.Println("Decremental:", x)*/

	//Retos
	//Area de un Rectangulo
	var rLargo int = 6
	var rAncho int = 2
	areaRectangulo := rLargo * rAncho

	fmt.Println("El Area del Rectangulo es :", areaRectangulo)

	//Area Trapecio
	base1T := 10
	base2T := 4
	alturaT := 4

	areaTrapecio := (base1T + base2T) * alturaT / 2
	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	//Area Circulo

	var areaC float64 = 4

	areaCirculo := math.Pi * math.Pow(areaC, 2)

	fmt.Println("El Area del Circulo es :", areaCirculo)
}
