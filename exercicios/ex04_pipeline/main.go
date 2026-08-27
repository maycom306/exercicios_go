package main

import (
	"fmt"
	"sync"
)

// ============================================================
//  EXERCÍCIO 4 — Pipeline Concorrente com Canais
// ============================================================
//
// OBJETIVO:
//   Entender o padrão pipeline: goroutines que se comunicam
//   via canais encadeados, cada uma responsável por um estágio.
//
// INSTRUÇÕES:
//
//  Implemente um pipeline de 3 estágios:
//
//    [Gerador] --canal1--> [Filtro] --canal2--> [Processador]
//
//  ESTÁGIO 1 — func gerador(max int) <-chan int
//    - Roda em uma goroutine
//    - Envia números de 1 até max para um canal
//    - Fecha o canal ao terminar
//    - Retorna o canal de saída (read-only)
//
//  ESTÁGIO 2 — func filtro(entrada <-chan int) <-chan int
//    - Roda em uma goroutine
//    - Lê do canal de entrada
//    - Deixa passar apenas os números PRIMOS
//    - Fecha o canal de saída ao terminar
//    - Retorna o canal de saída
//
//    DICA — Como verificar se um número é primo:
//      for i := 2; i*i <= n; i++ { if n%i == 0 { return false } }
//
//  ESTÁGIO 3 — func processador(entrada <-chan int, wg *sync.WaitGroup)
//    - Roda em uma goroutine
//    - Lê os primos do canal de entrada
//    - Para cada primo, imprime: "X² = Y"
//    - Chama wg.Done() ao terminar
//
//  No main():
//    - Encadeie os 3 estágios
//    - Use sync.WaitGroup para aguardar o processador terminar
//    - Passe max = 20
//
// SAÍDA ESPERADA:
//   2² = 4
//   3² = 9
//   5² = 25
//   7² = 49
//   11² = 121
//   ...
//
// ============================================================
func gerador(max int) <-chan int {
	canal := make(chan int)
	
	go func() {
		// O defer agenda o fechamento do canal para o final desta goroutine.
		// Assim que o loop 'for' acabar (ou se houver qualquer erro), o Go fecha o canal para nós!
		defer close(canal) 

		for i := 1; i <= max; i++ {
			canal <- i
		}
		
		// close(canal) -> ESSA LINHA NÃO É MAIS NECESSÁRIA AQUI!
	}()
	
	return canal
}

func primo(num int) bool{
	if num <= 1{
		return false
	}
	for i := 2; i*i <= num; i++{
		if num%i == 0{
			return false
		}
	}
	return true
}

func filtro(entrada <- chan int) <- chan int{
	saida := make(chan int)

	go func(){
		defer close(saida)
		
		for num := range entrada{
			if primo(num){
				saida <- num
			}
		}
	}()
	return saida 
}

func processar(entrada <-chan int, wg *sync.WaitGroup){
	defer wg.Done()

	for num := range entrada{
		fmt.Printf("%d² = %d\n", num, num*num)
	}
}

func main() {
	// Escreva seu código aqui
	var wg sync.WaitGroup
	var ate int
	fmt.Scanf("Digite um valor para saber os numeros primos ate ele e o seu quadrado: %d", &ate)
	wg.Add(1)
	go processar(filtro(gerador(ate)), &wg)
	wg.Wait()

}
