package main

import (
	"context"
	"math/rand"
	"sync"
	"time"
	"fmt"
)

// ============================================================
//  EXERCÍCIO 5 — Context com Timeout e Cancelamento
// ============================================================
//
// OBJETIVO:
//   Aprender a usar o pacote `context` para controlar o tempo
//   de vida de goroutines e cancelar operações em andamento.
//
// PACOTES NECESSÁRIOS: context, time, sync, math/rand
//
// INSTRUÇÕES:
//
//  1. No main(), crie um contexto com timeout de 1 segundo:
//       ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//       defer cancel()s
//
//  2. Impleemente a função:
//       func operacaoPesada(ctx contxt.Context, id int, wg *sync.WaitGroup, resultados chan<- Resultado)
//
//     Dentro dela:
//       a) Sorteie uma duração aleatória entre 100ms e 2000ms
//       b) Use um `select` com dois cases:
//            - case <-time.After(duracao): operação concluiu com sucesso
//            - case <-ctx.Done():          operação foi cancelada
//       c) Envie um struct Resultado para o canal com: ID, duracao e se concluiu
//       d) Imprima uma mensagem indicando se concluiu ou foi cancelada
//
//  3. Defina o struct Resultado:
//       type Resultado struct {
//           ID       int
//           Duracao  time.Duration
//           Concluiu bool
//       }
//
//  4. No main():
//       - Lance 5 goroutines de operacaoPesada
//       - Use sync.WaitGroup para saber quando todas terminaram
//       - Feche o canal de resultados após todas terminarem (use goroutine)
//       - Colete todos os resultados e imprima um resumo final:
//           "Concluídas: X | Canceladas: Y"
//
// CONCEITO-CHAVE:
//   O `select` em Go funciona como um switch para canais.
//   O case que receber um valor primeiro "vence" e é executado.
//   ctx.Done() retorna um canal que é fechado quando o contexto expira.
//
// ============================================================
type Resultado struct{
	ID int
	Duracao time.Duration
	Conclusao bool
}

func OperacaoPesada(ctx context.Context, ID int, wg *sync.WaitGroup, Resultados chan<- Resultado){
	duracao := time.Duration(rand.Intn(1901)+100) * time.Millisecond
	defer wg.Done()
	select {
	case <- time.After(duracao):
		res := Resultado{
			ID: ID,
			Duracao: duracao,
			Conclusao: true,
		}
		Resultados <- res
		fmt.Printf("Operação %d foi concluida com %s duração\n", res.ID, res.Duracao)
			
	case <-ctx.Done():
		res := Resultado{
			ID: ID,
			Duracao: duracao,
			Conclusao: false,
		}
		Resultados <- res
		fmt.Printf("[GOroutine %d] Cancelada! tempo que seria necessario: %s\n", res.ID, res.Duracao)
	}
	
}


func main() {
	// Escreva seu código aqui

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	resultados := make(chan Resultado, 5)
	var wg sync.WaitGroup

	fmt.Println("=== Iniciando 5 Goroutines Concorrentes com Timeout de 1s ===")
	
	for i := 1; i<=5; i++{
		wg.Add(1)
		go OperacaoPesada(ctx, i, &wg, resultados)
	}
	go func(){
		wg.Wait()
		close(resultados)
	}()
	concluidas := 0
	terminadas := 0

	for res := range resultados{
		if res.Conclusao == true{
			concluidas++
		}else{
			terminadas++
		}
	
	}
	fmt.Println("\n=== RESUMO FINAL ===")
	fmt.Printf("===== Concluidas [%d] ======= Caceladas [%d]\n", concluidas, terminadas)	
}
