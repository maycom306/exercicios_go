package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Contadores globais de testes
var (
	totalPassou int32
	totalFalhou int32
)

// TestMain executa antes e depois de toda a suíte de testes
func TestMain(m *testing.M) {
	// Executa os testes
	codigoSaida := m.Run()

	// Exibe o placar final detalhado
	fmt.Println("\n==============================================")
	fmt.Println("           RELATÓRIO FINAL DOS TESTES         ")
	fmt.Println("==============================================")
	fmt.Printf("  ✅ PASSOU:     %d\n", atomic.LoadInt32(&totalPassou))
	fmt.Printf("  ❌ NÃO PASSOU: %d\n", atomic.LoadInt32(&totalFalhou))
	fmt.Println("==============================================")

	os.Exit(codigoSaida)
}

// Registrar resultado auxiliar
func registrarResultado(t *testing.T, passou bool, msgSucesso, msgFalha string) {
	t.Helper()
	if passou {
		atomic.AddInt32(&totalPassou, 1)
		t.Logf("✅ PASSOU | %s", msgSucesso)
	} else {
		atomic.AddInt32(&totalFalhou, 1)
		t.Errorf("❌ NÃO PASSOU | %s", msgFalha)
	}
}

// ============================================================
// TESTES UNITÁRIOS — Algoritmo de Números Primos
// ============================================================
func TestPrimo(t *testing.T) {
	casosDeTeste := []struct {
		nome     string
		numero   int
		esperado bool
	}{
		{nome: "Número negativo (-5)", numero: -5, esperado: false},
		{nome: "Zero (0)", numero: 0, esperado: false},
		{nome: "Número Um (1)", numero: 1, esperado: false},
		{nome: "Número Dois (2)", numero: 2, esperado: true},
		{nome: "Número Três (3)", numero: 3, esperado: true},
		{nome: "Número Quatro (4)", numero: 4, esperado: false},
		{nome: "Número Treze (13)", numero: 13, esperado: true},
		{nome: "Número Quinze (15)", numero: 15, esperado: false},
		{nome: "Número Vinte e Cinco (25)", numero: 25, esperado: false},
	}

	for _, tt := range casosDeTeste {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := primo(tt.numero)
			passou := (resultado == tt.esperado)

			msgSucesso := fmt.Sprintf("Entrada: %d | Retorno: %v | Esperado: %v", tt.numero, resultado, tt.esperado)
			msgFalha := fmt.Sprintf("Entrada: %d | Retorno: %v | Esperado: %v", tt.numero, resultado, tt.esperado)

			registrarResultado(t, passou, msgSucesso, msgFalha)
		})
	}
}

// ============================================================
// TESTES UNITÁRIOS — Estágio 1: Gerador
// ============================================================
func TestEstagioGerador(t *testing.T) {
	t.Run("Validar sequência gerada de 1 até MAX", func(t *testing.T) {
		maximo := 5
		canal := gerador(maximo)

		var recebidos []int
		for val := range canal {
			recebidos = append(recebidos, val)
		}

		esperado := []int{1, 2, 3, 4, 5}
		passou := reflect.DeepEqual(recebidos, esperado)

		msgSucesso := fmt.Sprintf("gerador(%d) produziu a sequência %v", maximo, recebidos)
		msgFalha := fmt.Sprintf("gerador(%d) produziu %v, esperava %v", maximo, recebidos, esperado)

		registrarResultado(t, passou, msgSucesso, msgFalha)
	})
}

// ============================================================
// TESTES UNITÁRIOS — Estágio 2: Filtro
// ============================================================
func TestEstagioFiltro(t *testing.T) {
	t.Run("Validar se deixa passar apenas primos", func(t *testing.T) {
		entrada := make(chan int, 10)
		dados := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		for _, v := range dados {
			entrada <- v
		}
		close(entrada)

		canalSaida := filtro(entrada)

		var primosFiltrados []int
		for p := range canalSaida {
			primosFiltrados = append(primosFiltrados, p)
		}

		esperado := []int{2, 3, 5, 7}
		passou := reflect.DeepEqual(primosFiltrados, esperado)

		msgSucesso := fmt.Sprintf("Entrada %v -> Filtro produziu %v", dados, primosFiltrados)
		msgFalha := fmt.Sprintf("Entrada %v -> Filtro produziu %v, esperava %v", dados, primosFiltrados, esperado)

		registrarResultado(t, passou, msgSucesso, msgFalha)
	})
}

// ============================================================
// TESTES UNITÁRIOS — Estágio 3: Processador
// ============================================================
func TestEstagioProcessador(t *testing.T) {
	t.Run("Validar formatação da saída impressa no console", func(t *testing.T) {
		entrada := make(chan int, 3)
		entrada <- 2
		entrada <- 3
		entrada <- 5
		close(entrada)

		var wg sync.WaitGroup
		wg.Add(1)

		r, w, _ := os.Pipe()
		stdoutOriginal := os.Stdout
		os.Stdout = w

		processar(entrada, &wg)
		wg.Wait()

		w.Close()
		os.Stdout = stdoutOriginal

		var buf bytes.Buffer
		io.Copy(&buf, r)
		saidaTexto := strings.TrimSpace(buf.String())

		saidaEsperada := "2² = 4\n3² = 9\n5² = 25"
		passou := (saidaTexto == saidaEsperada)

		msgSucesso := fmt.Sprintf("Formatação gerada com sucesso:\n%s", saidaTexto)
		msgFalha := fmt.Sprintf("Obtido:\n%q\nEsperado:\n%q", saidaTexto, saidaEsperada)

		registrarResultado(t, passou, msgSucesso, msgFalha)
	})
}

// ============================================================
// TESTES DE INTEGRAÇÃO — Pipeline Completo (max = 20)
// ============================================================
func TestPipelineCompleto(t *testing.T) {
	t.Run("Execução encadeada do exercício (max = 20)", func(t *testing.T) {
		maximo := 20
		var wg sync.WaitGroup
		wg.Add(1)

		r, w, _ := os.Pipe()
		stdoutOriginal := os.Stdout
		os.Stdout = w

		go processar(filtro(gerador(maximo)), &wg)
		wg.Wait()

		w.Close()
		os.Stdout = stdoutOriginal

		var buf bytes.Buffer
		io.Copy(&buf, r)
		linhasObtidas := strings.Split(strings.TrimSpace(buf.String()), "\n")

		linhasEsperadas := []string{
			"2² = 4",
			"3² = 9",
			"5² = 25",
			"7² = 49",
			"11² = 121",
			"13² = 169",
			"17² = 289",
			"19² = 361",
		}

		passou := reflect.DeepEqual(linhasObtidas, linhasEsperadas)

		msgSucesso := fmt.Sprintf("Pipeline gerou todas as %d linhas esperadas", len(linhasObtidas))
		msgFalha := fmt.Sprintf("Resultado do Pipeline incorreto.\nObtido: %v\nEsperado: %v", linhasObtidas, linhasEsperadas)

		registrarResultado(t, passou, msgSucesso, msgFalha)
	})
}

// ============================================================
// TESTES DE CONCORRÊNCIA — Checagem de Deadlock
// ============================================================
func TestConcorrenciaEDeadlock(t *testing.T) {
	t.Run("Garantir encerramento de goroutines e canais sem travar", func(t *testing.T) {
		canalSucesso := make(chan bool)

		go func() {
			var wg sync.WaitGroup
			wg.Add(1)

			nullOut, _ := os.Open(os.DevNull)
			stdoutOriginal := os.Stdout
			os.Stdout = nullOut

			go processar(filtro(gerador(50)), &wg)
			wg.Wait()

			os.Stdout = stdoutOriginal
			canalSucesso <- true
		}()

		select {
		case <-canalSucesso:
			registrarResultado(t, true, "Canais e goroutines fecharam sem deadlock.", "")
		case <-time.After(2 * time.Second):
			registrarResultado(t, false, "", "DEADLOCK DETECTADO: O pipeline estourou o limite de 2 segundos.")
		}
	})
}