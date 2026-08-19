package main

// ============================================================
//  EXERCÍCIO 7 — Middleware de Autenticação com Chi
// ============================================================
//
// OBJETIVO:
//   Aprender a criar e encadear middlewares HTTP com chi,
//   incluindo autenticação via API Key e logging customizado.
//
// SETUP:
//   go mod init ex07_middleware
//   go get github.com/go-chi/chi/v5
//
// INSTRUÇÕES:
//
//  1. Crie um middleware `RequireAPIKey` que:
//       - Lê o header "X-API-Key" da requisição
//       - Se o header estiver ausente ou incorreto → responde 401 Unauthorized
//       - Se estiver correto → chama next.ServeHTTP(w, r) para continuar
//       - A chave válida pode ser uma constante: const apiKey = "minha-chave-secreta"
//
//     Assinatura de um middleware chi:
//       func RequireAPIKey(next http.Handler) http.Handler {
//           return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//               // sua lógica aqui
//           })
//       }
//
//  2. Crie um middleware `Logger` customizado que:
//       - Registra: método HTTP, path, e tempo de execução do handler
//       - Use time.Now() antes e time.Since() depois de chamar next
//       - Imprima no formato: "[GET] /rota → 200 em 1.2ms"
//
//  3. Crie um middleware `CORS` simples que:
//       - Adiciona o header "Access-Control-Allow-Origin: *"
//       - Responde 200 em requisições OPTIONS (preflight)
//
//  4. Monte o router com dois grupos de rotas:
//
//     Rotas PÚBLICAS (sem autenticação):
//       GET /health → retorna {"status": "ok"} em JSON
//       GET /info   → retorna {"app": "ex07", "versao": "1.0"}
//
//     Rotas PROTEGIDAS (exigem X-API-Key):
//       GET  /admin/usuarios    → retorna lista de usuários fake
//       POST /admin/usuarios    → simula criação de usuário
//
//  5. Aplique os middlewares Logger e CORS globalmente,
//     e RequireAPIKey apenas no grupo /admin.
//
// TESTE:
//   # Deve funcionar (pública):
//   curl http://localhost:8080/health
//
//   # Deve retornar 401:
//   curl http://localhost:8080/admin/usuarios
//
//   # Deve funcionar (com chave correta):
//   curl -H "X-API-Key: minha-chave-secreta" http://localhost:8080/admin/usuarios
//
// ============================================================

func main() {
	// Escreva seu código aqui
}
