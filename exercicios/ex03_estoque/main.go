package main

// ============================================================
//  EXERCÍCIO 3 — Gerenciador de Estoque (maps + error handling)
// ============================================================
//
// OBJETIVO:
//   Praticar maps, error handling com o tipo `error` nativo,
//   ponteiros e organização de código em Go.
//
// INSTRUÇÕES:
//
//  1. Crie a struct `Produto` com os campos:
//       - ID         int
//       - Nome       string
//       - Preco      float64
//       - Quantidade int
//
//  2. Crie a struct `Estoque` com:
//       - Um map[int]Produto como "banco de dados" em memória
//       - Um campo para controlar o próximo ID automaticamente
//
//  3. Implemente os seguintes métodos em *Estoque:
//
//     a) AdicionarProduto(nome string, preco float64, qtd int) Produto
//          → Cria e armazena um produto com ID auto-incrementado
//
//     b) RemoverEstoque(id, quantidade int) error
//          → Retorna erro se: produto não existe OU quantidade insuficiente
//          → Se ok, subtrai a quantidade do estoque
//
//     c) BuscarPorNome(nome string) ([]Produto, error)
//          → Busca case-insensitive (use strings.ToLower)
//          → Retorna erro se nenhum produto for encontrado
//
//     d) TotalEmEstoque() (totalItens int, totalValor float64)
//          → Retorna o total de itens e o valor total em estoque
//          → Use retorno nomeado (named return values)
//
//  4. No main(), demonstre cada função com prints explicativos.
//     Teste também os casos de erro (estoque insuficiente, produto inexistente).
//
// DICA — Como criar e retornar errors:
//   import "errors"
//   return errors.New("mensagem de erro")
//   return fmt.Errorf("erro: produto %d não encontrado", id)
//
// ============================================================

func main() {
	// Escreva seu código aqui
}
