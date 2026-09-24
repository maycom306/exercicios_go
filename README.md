# 🐹 Exercícios de Go

Repositório de exercícios práticos para aprendizado da linguagem **Go**, cobrindo desde conceitos fundamentais até tópicos avançados como concorrência e desenvolvimento de APIs REST.

---

## 📁 Estrutura do Projeto

```
exercicios/
├── ex01_imc/        → Structs e Métodos
├── ex02_hof/        → Higher-Order Functions
├── ex03_estoque/    → Maps e Error Handling
├── ex04_pipeline/   → Pipeline Concorrente com Canais
├── ex05_context/    → Context com Timeout e Cancelamento
├── ex06_Fiber_api/    → API REST com Fiber
├── ex07_middleware/ → Middleware de Autenticação com Chi
├── ex08_cobra_cli/  → CLI Tool com Cobra
├── ex09_cache/      → Cache Concorrente com sync.RWMutex
└── ex10_gin_api/    → API Completa com Gin
```

---

## 🧪 Exercícios

### Ex 01 — Calculadora de IMC com Structs e Métodos
> **Conceitos:** `struct`, métodos, `switch`

Cria uma struct `Pessoa` com campos `Nome`, `Peso` e `Altura`. Implementa os métodos `IMC()` e `Classificacao()` e itera sobre uma slice de pessoas exibindo o resultado de cada uma.

```
Ana     | IMC: 22.50 | Normal
Carlos  | IMC: 30.12 | Obesidade
```

---

### Ex 02 — Higher-Order Functions: Filtrar / Mapear / Reduzir
> **Conceitos:** funções como parâmetros, first-class functions

Implementa as funções genéricas `Filtrar`, `Mapear` e `Reduzir` que operam sobre `[]int`. Aplica as três em cadeia para filtrar pares, dobrar os valores e somar o resultado.

```
Pares:    [2 4 6 8 10]
Dobrados: [4 8 12 16 20]
Soma:     60
```

---

### Ex 03 — Gerenciador de Estoque
> **Conceitos:** `map`, error handling, ponteiros, retorno nomeado

Implementa uma struct `Estoque` com operações de CRUD em memória: `AdicionarProduto`, `RemoverEstoque`, `BuscarPorNome` (case-insensitive) e `TotalEmEstoque`. Inclui tratamento de erros com `errors.New` e `fmt.Errorf`.

---

### Ex 04 — Pipeline Concorrente com Canais
> **Conceitos:** goroutines, channels, `sync.WaitGroup`, padrão pipeline

Constrói um pipeline de 3 estágios encadeados via canais: **Gerador** → **Filtro de Primos** → **Processador**. O processador imprime o quadrado de cada número primo recebido.

```
2² = 4
3² = 9
5² = 25
...
```

---

### Ex 05 — Context com Timeout e Cancelamento
> **Conceitos:** `context.WithTimeout`, `select`, goroutines, channels

Lança 5 goroutines simulando operações pesadas com duração aleatória (100ms a 2s) em um contexto com timeout de 1 segundo. Usa `select` para detectar conclusão ou cancelamento e imprime um resumo final.

```
Concluídas: 3 | Canceladas: 2
```

---

### Ex 06 — API REST com Fiber
> **Conceitos:** Fiber, handlers HTTP, JSON, `sync.RWMutex`

Cria uma API REST de produtos com banco de dados em memória. Implementa os endpoints `GET /produtos`, `GET /produtos/{id}`, `POST /produtos` e `DELETE /produtos/{id}`.

**Setup:**
```bash
go mod init ex06_fiber
go get github.com/gofiber/fiber/v2
go run main.go
```

---

### Ex 07 — Middleware de Autenticação com Chi
> **Conceitos:** middlewares HTTP, autenticação por API Key, CORS, logging

Implementa três middlewares do zero: `RequireAPIKey` (autenticação por header), `Logger` (método, path e tempo de execução) e `CORS`. Organiza rotas em grupos públicos e protegidos.

```bash
# Público
curl http://localhost:8080/health

# Protegido (retorna 401 sem a chave)
curl -H "X-API-Key: minha-chave-secreta" http://localhost:8080/admin/usuarios
```

---

### Ex 08 — CLI Tool com Cobra
> **Conceitos:** `spf13/cobra`, subcomandos, flags, calculadora financeira

Cria a CLI `calcfin` com três subcomandos:
- `imc` — calcula IMC e classificação a partir de `--peso` e `--altura`
- `salario` — calcula salário líquido com desconto de INSS progressivo
- `juros` — calcula juros compostos mês a mês com `--capital`, `--taxa` e `--meses`

```bash
go run main.go imc --peso 70 --altura 1.75
# → IMC: 22.86 | Normal

go run main.go salario --bruto 3500
# → Bruto: R$ 3500.00 | Desconto: R$ 420.00 | Líquido: R$ 3080.00
```

---

### Ex 09 — Cache Concorrente com sync.RWMutex
> **Conceitos:** `sync.RWMutex`, concorrência segura, race detector

Implementa uma struct `Cache` thread-safe com os métodos `Set`, `Get`, `Delete` e `Len`. Testa a concorrência com 5 goroutines de escrita e 10 de leitura simultâneas.

```bash
# Verificar ausência de race conditions
go run -race main.go
```

---

### Ex 10 — API Completa com Gin + Middleware + Validação
> **Conceitos:** `gin-gonic/gin`, validação de input, middlewares, grupos de rotas

API de gerenciamento de tarefas (Todo API) com:
- Validação automática via tags `binding:"required,min=3"`
- Endpoints CRUD completos em `/api/v1/tarefas`
- Middleware de logging customizado e de Request ID
- Respostas de erro padronizadas

```bash
go mod init ex10_gin
go get github.com/gin-gonic/gin
go run main.go
```

---

## 🛠️ Pré-requisitos

- [Go 1.21+](https://go.dev/dl/)

## 🚀 Como executar um exercício

```bash
# Navegue até o diretório do exercício
cd exercicios/ex01_imc

# Execute
go run main.go
```

> Para exercícios com dependências externas (ex06, ex07, ex08, ex10), siga as instruções de **Setup** descritas dentro do próprio `main.go` de cada exercício.
