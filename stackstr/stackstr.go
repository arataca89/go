// Implementa uma pilha de strings
// arataca89@gmail.com
// 20210622

package stackstr

//package main

import (
	"fmt"
)

// StackStr é uma estrutura que define uma pilha de strings
type StackStr []string

// isVoid verifica se a pilha está vazia
func (p StackStr) IsVoid() bool {
	return len(p) == 0
}

// push insere um item na pilha
func (p *StackStr) Push(s string) {
	*p = append(*p, s)
}

// pop retira um item da pilha
func (p *StackStr) Pop() string {
	s := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return s
}

// getTop retorna o item do topo da pilha
func (p StackStr) GetTop() string {
	return p[len(p)-1]
}

// printStack imprime os dados da pilha
func (p StackStr) PrintStack() {
	fmt.Println("Top  :", len(p))
	fmt.Println("Pilha:", p)
}

// TESTES
//
// func main() {
// 	fmt.Printf("\n>>> stackstr <<<\n")
// 	var pilha StackStr
// 	expressao := "(13 * 41)=(11/85)"
// 	for i := 0; i < len(expressao); i++ {
// 		pilha.Push(string(expressao[i]))
// 	}
// 	pilha.PrintStack()

// 	// testes
// 	fmt.Println()
// 	fmt.Println("top():", pilha.GetTop())
// 	fmt.Println("pop():", pilha.Pop())
// 	pilha.PrintStack()

// 	fmt.Println()
// 	fmt.Println("top():", pilha.GetTop())
// 	fmt.Println("pop():", pilha.Pop())
// 	pilha.PrintStack()

// 	fmt.Println()
// 	fmt.Println("top():", pilha.GetTop())
// 	fmt.Println("pop():", pilha.Pop())
// 	pilha.PrintStack()
// }
