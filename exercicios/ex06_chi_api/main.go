package main

// ============================================================
//  EXERCÍCIO 6 — API REST com Router Chi
// ============================================================
//
// OBJETIVO:
//   Criar uma API REST usando o router `chi`, um dos mais
//   populares e idiomáticos roteadores HTTP do ecossistema Go.
//
// SETUP (rode antes de começar):
//   go mod init ex06_chi
//   go get github.com/go-chi/chi/v5
//
// INSTRUÇÕES:
//
//  1. Crie a struct `Produto` com campos:
//       ID    int    `json:"id"`
//       Nome  string `json:"nome"`
//       Preco float64 `json:"preco"`
//
//  2. Crie um "banco de dados" em memória:
//       var produtos = map[int]Produto{}
//       var proximoID = 1
//       var mu sync.RWMutex   // para proteger acesso concorrente
//
//  3. Implemente os seguintes handlers HTTP:
//
//     GET  /produtos
//       → Retorna todos os produtos em JSON
//       → Use mu.RLock() / mu.RUnlock() para leitura segura
//
//     GET  /produtos/{id}
//       → Lê o parâmetro com: chi.URLParam(r, "id")
//       → Converta para int com strconv.Atoi
//       → Retorna 404 se não encontrado
//
//     POST /produtos
//       → Lê o body com json.NewDecoder(r.Body).Decode(&p)
//       → Atribui um ID automático
//       → Retorna 201 Created com o produto criado em JSON
//
//     DELETE /produtos/{id}
//       → Remove o produto pelo ID
//       → Retorna 204 No Content se removido, 404 se não encontrado
//
//  4. No main(), configure o router chi e suba o servidor:
//       r := chi.NewRouter()
//       r.Use(middleware.Logger)     // log de requisições
//       r.Use(middleware.Recoverer)  // recupera panics
//       // registre as rotas aqui
//       http.ListenAndServe(":8080", r)
//
// DICA — Retornar JSON:
//   w.Header().Set("Content-Type", "application/json")
//   json.NewEncoder(w).Encode(dados)
//
// TESTE com curl ou Postman:
//   curl http://localhost:8080/produtos
//   curl -X POST http://localhost:8080/produtos \
//        -H "Content-Type: application/json" \
//        -d '{"nome":"Notebook","preco":3500.00}'
//   curl http://localhost:8080/produtos/1
//   curl -X DELETE http://localhost:8080/produtos/1
//
// ============================================================

func main() {
	// Escreva seu código aqui
}
