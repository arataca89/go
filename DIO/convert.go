/*
convert.go
Converte o ponto de ebulicao de graus Kelvin para graus Celsius
C = K - 273
*/

package main

import "fmt"

func main(){
	kelvin := 373;
	celsius := kelvin - 273;
	fmt.Println("Ponto de ebulicao em Kelvin : ",kelvin)
	fmt.Println("Ponto de ebulicao em Celsius: ",celsius)
}
