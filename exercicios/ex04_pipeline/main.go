package main

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

func main() {
	// Escreva seu código aqui
}
