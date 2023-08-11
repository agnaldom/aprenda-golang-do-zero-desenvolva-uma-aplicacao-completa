package main

import "fmt"

func main() {
	var numero int64 = -1000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	// FIM NUMEROS INTEIROS

	// NUMEROS REAIS

	var numeroReal1 float32 = 123.4
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000.34
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// FIM NUMERO REAIS

}
