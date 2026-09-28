package main

import (
	"fmt"
	"sync"
	"time"
)

type Tarefa struct {
	ID    int
	Bruto float64
}

func worker(workerID int, tarefas <-chan Tarefa, wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range tarefas {
		var desconto float64
		switch {
		case t.Bruto <= 1621.00:
			desconto = t.Bruto * 0.075
		case t.Bruto <= 2902.84:
			desconto = t.Bruto * 0.09
		case t.Bruto <= 4354.27:
			desconto = t.Bruto * 0.12
		default:
			desconto = t.Bruto * 0.14
		}

		liquido := t.Bruto - desconto
		fmt.Printf("[Worker %d] Tarefa #%d | Bruto: R$ %.2f | Líquido: R$ %.2f\n", 
			workerID, t.ID, t.Bruto, liquido)

		// Simula um processamento de 500ms para visualizar os 3 workers trabalhando
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	salarios := []float64{1412.00, 2200.50, 3500.00, 4354.27, 6800.00, 12000.00, 2800.00, 9500.00}

	const maxWorkers = 3
	canalTarefas := make(chan Tarefa, len(salarios))
	var wg sync.WaitGroup

	// 1. Inicia exatamente 3 workers
	for w := 1; w <= maxWorkers; w++ {
		wg.Add(1)
		go worker(w, canalTarefas, &wg)
	}

	// 2. Envia os salários para a fila
	for i, bruto := range salarios {
		canalTarefas <- Tarefa{ID: i + 1, Bruto: bruto}
	}
	close(canalTarefas) // Fecha o canal para avisar os workers que não há mais tarefas

	// 3. Aguarda os 3 workers terminarem tudo
	wg.Wait()
	fmt.Println("\nTodos os cálculos foram finalizados!")
}