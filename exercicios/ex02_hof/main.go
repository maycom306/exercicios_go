package main

import "fmt"

// ============================================================
//  EXERCÍCIO 2 — Higher-Order Functions: Filtrar / Mapear / Reduzir
// ============================================================
//
// OBJETIVO:
//   Entender funções como parâmetros (first-class functions) em Go.
//
// INSTRUÇÕES:
//
//  1. Implemente a função `Filtrar`:
//       - Recebe: slice []int  e  condicao func(int) bool
//       - Retorna: []int com apenas os elementos que passarem na condição
//
//  2. Implemente a função `Mapear`:
//       - Recebe: slice []int  e  transformar func(int) int
//       - Retorna: []int com a transformação aplicada em cada elemento
//
//  3. Implemente a função `Reduzir`:
//       - Recebe: slice []int, acumulador func(int, int) int, valorInicial int
//       - Retorna: int com o resultado acumulado de todos os elementos
//
//  4. No main(), use as três funções sobre []int{1,2,3,4,5,6,7,8,9,10} para:
//       a) Filtrar apenas os números PARES
//       b) DOBRAR cada número filtrado
//       c) SOMAR todos eles
//       d) Imprimir o resultado final
//
// SAÍDA ESPERADA:
//   Pares:    [2 4 6 8 10]
//   Dobrados: [4 8 12 16 20]
//   Soma:     60
//
// DESAFIO EXTRA:
//   Encadeie as três chamadas em UMA ÚNICA expressão dentro do main().
//
// ============================================================
func filtrar(num[]int,condicao func(int) bool)[]int{	
	var resultado[]int
	for _, valor := range num{
		if condicao (valor){
			resultado = append(resultado, valor)
		}
	}
	return resultado
}

func mapear(num[]int, tranforma func(int)int)[]int{
	var resultado []int
	for _, valor := range num{
		novoValor := tranforma(valor)
		resultado = append(resultado, novoValor)
	}
	return resultado
}

func reduzir(num[]int, acumular func(int, int)int, valorInicial int) int{
	total := valorInicial
	for _, valor := range num{
		total = acumular(total, valor)

	}
	return total
}

func main() {
	numero := []int{1,2,3,4,5,6,7,8,9,10}// Escreva seu código aqui

	par := filtrar(numero, func (n int)bool{
		return n%2	== 0
	})

	dobro := mapear(par, func (n int) int{
		return n*2 
	})
	soma := reduzir(dobro, func (acumular, valorInicial int) int{
		return acumular + valorInicial
	}, 0)
	
	fmt.Println("Original:", numero)
	fmt.Println("Pares:", par)
	fmt.Println("Dobrados:", dobro)
	fmt.Println("Soma total:", soma)
}
