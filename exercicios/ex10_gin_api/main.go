package main

// ============================================================
//  EXERCÍCIO 10 — API Completa com Gin + Middleware + Validação
// ============================================================
//
// OBJETIVO:
//   Construir uma API REST profissional usando o framework Gin,
//   com validação de entrada, middleware, e tratamento de erros.
//
// SETUP:
//   go mod init ex10_gin
//   go get github.com/gin-gonic/gin
//
// CONTEXTO:
//   Gin é um framework HTTP de alta performance muito usado
//   em produção. Ele tem binding automático de JSON, validação,
//   grupos de rotas e middlewares embutidos.
//
// INSTRUÇÕES:
//
//  Implemente uma API de gerenciamento de tarefas (Todo API).
//
//  ─────────────────────────────────────────────────
//  STRUCT E VALIDAÇÃO
//  ─────────────────────────────────────────────────
//  Crie a struct `Tarefa` com tags de binding/validação:
//
//    type Tarefa struct {
//        ID        int    `json:"id"`
//        Titulo    string `json:"titulo"    binding:"required,min=3"`
//        Descricao string `json:"descricao"`
//        Feita     bool   `json:"feita"`
//    }
//
//  A tag `binding:"required,min=3"` faz o Gin validar automaticamente.
//
//  ─────────────────────────────────────────────────
//  ENDPOINTS
//  ─────────────────────────────────────────────────
//
//  GET    /api/v1/tarefas          → lista todas
//  GET    /api/v1/tarefas/:id      → busca por ID → 404 se não existir
//  POST   /api/v1/tarefas          → cria nova (valida o body)
//                                  → 400 se inválido
//  PUT    /api/v1/tarefas/:id      → atualiza tarefa completa
//  PATCH  /api/v1/tarefas/:id/done → marca como feita (toggle)
//  DELETE /api/v1/tarefas/:id      → remove → 204 No Content
//
//  ─────────────────────────────────────────────────
//  MIDDLEWARES (implemente como funções gin.HandlerFunc)
//  ─────────────────────────────────────────────────
//
//  1. LoggerMiddleware:
//       Imprime: [MÉTODO] /path | status | tempo
//       Use c.Next() para executar o handler antes de logar o status.
//       Dica: status = c.Writer.Status()
//
//  2. RequestIDMiddleware:
//       Gera um ID único por requisição (pode ser um contador ou UUID simples)
//       Adiciona ao header de resposta: "X-Request-ID: <id>"
//       Salva no contexto Gin: c.Set("requestID", id)
//
//  ─────────────────────────────────────────────────
//  RESPOSTAS PADRONIZADAS
//  ─────────────────────────────────────────────────
//  Crie um helper para respostas de erro padronizadas:
//
//    type ErrResponse struct {
//        Erro    string `json:"erro"`
//        Detalhe string `json:"detalhe,omitempty"`
//    }
//
//    func respondErro(c *gin.Context, status int, msg string) {
//        c.JSON(status, ErrResponse{Erro: msg})
//    }
//
//  ─────────────────────────────────────────────────
//  MAIN
//  ─────────────────────────────────────────────────
//  Configure o Gin:
//    r := gin.New()                  // sem o logger padrão
//    r.Use(LoggerMiddleware)         // seu logger customizado
//    r.Use(RequestIDMiddleware)
//    r.Use(gin.Recovery())           // recupera panics
//
//    api := r.Group("/api/v1")
//    // registre as rotas aqui
//
//    r.Run(":8080")
//
//  ─────────────────────────────────────────────────
//  TESTE:
//  ─────────────────────────────────────────────────
//  # Criar tarefa válida:
//  curl -X POST http://localhost:8080/api/v1/tarefas \
//       -H "Content-Type: application/json" \
//       -d '{"titulo":"Estudar Go","descricao":"Praticar exercicios"}'
//
//  # Criar tarefa inválida (título muito curto):
//  curl -X POST http://localhost:8080/api/v1/tarefas \
//       -H "Content-Type: application/json" \
//       -d '{"titulo":"Go"}'
//
//  # Listar:
//  curl http://localhost:8080/api/v1/tarefas
//
//  # Marcar como feita:
//  curl -X PATCH http://localhost:8080/api/v1/tarefas/1/done
//
// ============================================================

func main() {
	// Escreva seu código aqui
}
