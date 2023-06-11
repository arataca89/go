/********************************************************************
numeros1.go

Exibe todos os numeros entre 1 e 100 que sao divisiveis por 3

*********************************************************************/

package main

import "fmt"

func main(){
	for i := 1; i <= 100; i++{
		if i % 3 == 0{
			fmt.Print(i," ")
		}
	}
}
/********************************************************************
SAIDA:

C:\Users\Documentos\go\src\desafios_dio>go run numeros1.go
3 6 9 12 15 18 21 24 27 30 33 36 39 42 45 48 51 54 57 60 63 66 69 72 75 78 81 84 87 90 93 96 99
*********************************************************************/
