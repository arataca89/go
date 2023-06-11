/********************************************************************
numeros2.go

Exibe todos os numeros entre 1 e 100.
Se o numero for divisivel por 3 o programa escreve "Pin"
Se o numero for divisivel por 5 o programa escreve "Pan"
Se o numero for divisivel por 3 e por 5 o programa escreve "PinPan"

*********************************************************************/

package main

import "fmt"

func novalinha(n int){
	if n % 10 == 0{
		fmt.Println()
	}
}

func main(){
	for i := 1; i <= 100; i++{ 		
		if i % 3 == 0 && i % 5 == 0{
			fmt.Print("PinPan ")
			novalinha(i)
			continue
		}
		if i % 3 == 0{
			fmt.Print("Pin ")
			novalinha(i)
			continue
		}
		if i % 5 == 0{
			fmt.Print("Pan ")
			novalinha(i)
			continue
		}
		fmt.Print(i," ")
		novalinha(i)
	}
}
/********************************************************************
SAIDA:

C:\Users\Documentos\go\src\desafios_dio>go run numeros2.go
1 2 Pin 4 Pan Pin 7 8 Pin Pan
11 Pin 13 14 PinPan 16 17 Pin 19 Pan
Pin 22 23 Pin Pan 26 Pin 28 29 PinPan
31 32 Pin 34 Pan Pin 37 38 Pin Pan
41 Pin 43 44 PinPan 46 47 Pin 49 Pan
Pin 52 53 Pin Pan 56 Pin 58 59 PinPan
61 62 Pin 64 Pan Pin 67 68 Pin Pan
71 Pin 73 74 PinPan 76 77 Pin 79 Pan
Pin 82 83 Pin Pan 86 Pin 88 89 PinPan
91 92 Pin 94 Pan Pin 97 98 Pin Pan

*********************************************************************/
