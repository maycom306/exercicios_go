package main

// ============================================================
//  EXERCÍCIO 9 — Cache Concorrente com sync.RWMutex
// ============================================================
//
// OBJETIVO:
//   Entender a diferença entre Mutex e RWMutex, e como proteger
//   estruturas de dados compartilhadas em ambientes concorrentes.
//
// PACOTES NECESSÁRIOS: sync, fmt, time
//
// CONTEXTO:
//   sync.RWMutex permite:
//     - Múltiplos leitores simultâneos (RLock / RUnlock)
//     - Apenas UM escritor por vez, com exclusão total (Lock / Unlock)
//   Use quando leituras são muito mais frequentes que escritas.
//
// INSTRUÇÕES:
//
//  1. Crie a struct `Cache` com:
//       mu    sync.RWMutex
//       dados map[string]string
//
//  2. Implemente os seguintes métodos em *Cache:
//
//     Set(chave, valor string)
//       → Bloqueia para escrita (mu.Lock / mu.Unlock)
//       → Armazena o valor no map
//
//     Get(chave string) (string, bool)
//       → Bloqueia para leitura (mu.RLock / mu.RUnlock)
//       → Retorna o valor e true se existir, "" e false se não
//
//     Delete(chave string)
//       → Bloqueia para escrita
//       → Remove a chave do map
//
//     Len() int
//       → Bloqueia para leitura
//       → Retorna o número de entradas no cache
//
//  3. No main(), teste a concorrência:
//       a) Lance 5 goroutines de ESCRITA, cada uma fazendo
//          10 chamadas a Set() com chaves como "key-0", "key-1", etc.
//
//       b) Lance 10 goroutines de LEITURA, cada uma fazendo
//          20 chamadas a Get() para chaves aleatórias
//
//       c) Use sync.WaitGroup para aguardar todas terminarem
//
//       d) No final, imprima: "Cache contém X entradas"
//
//  4. DESAFIO: Rode com o detector de race conditions:
//       go run -race main.go
//     Seu código deve rodar SEM erros de race condition.
//
// DICA — Estrutura básica:
//   type Cache struct {
//       mu    sync.RWMutex
//       dados map[string]string
//   }
//
//   func NovoCache() *Cache {
//       return &Cache{dados: make(map[string]string)}
//   }
//
// ============================================================

func main() {
	// Escreva seu código aqui
}
