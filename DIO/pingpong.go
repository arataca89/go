/********************************************************************
pingpong.go

Desafio DIO:
	Programa que exiba na tela em alternancia as palavras "ping" e
	"pong".
	O programa deve usar canais e goroutines.
********************************************************************/
package main
import(
	"fmt"	
	"time"
)

func ping(c chan string){
	c <- "ping"	
}

func pong(c chan string){
	c <- "pong"	
}

func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	
	for i:=0;i<10;i++{
		go ping(c1);
		go pong(c2);
		time.Sleep(1 * time.Second)
		select{
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
		}
	}
}
/********************************************************************
SAIDA
=====

C:\Users\Documentos\go\src\dio\desafios_dio>go run pingpong.go
pong
ping
ping
pong
pong
ping
pong
ping
pong
pong

C:\Users\Documentos\go\src\dio\desafios_dio>go run pingpong.go
ping
ping
pong
ping
pong
pong
pong
pong
pong
ping

C:\Users\Documentos\go\src\dio\desafios_dio>go run pingpong.go

********************************************************************/
