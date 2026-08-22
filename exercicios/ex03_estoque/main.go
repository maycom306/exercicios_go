package main

import (
	"errors"
	"strings"
)

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
//     a) AdicionarProduto(nome string, preco float64, Qnt int) Produto
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

type Produto struct {
	ID int
	Nome string
	Preco float64
	Qnt int
}

type Estoque struct{
	BancoDados map[int]Produto
	ItemID int
}
func (e *Estoque) AdcionarProduto (Nome string, Preco float64, Qnt int) Produto{
	NovoItem := Produto{
		Nome : Nome,
		ID : e.ItemID,
		Preco : Preco,
		Qnt : Qnt,
	}
	e.BancoDados[e.ItemID] = NovoItem
	e.ItemID++
	return NovoItem 
}

func (r *Estoque)RemoverEstoque(ID int, Qnt int) error{
	Existente, existe := r.BancoDados[ID]
	if !existe {
		return errors.New("item Não encontrado")
	}
	if  Qnt > Existente.Qnt{
		return errors.New("Estoque Insuficiente")
	}
	Existente.Qnt -= Qnt
	r.BancoDados[ID] = Existente
	return nil
}

func (r *Estoque) BuscarPorNome(Nome string) ([]Produto,error){
	var ItensEncontrados []Produto
	minusculo := strings.ToLower(Nome)
	for _, Produto := range r.BancoDados{
		if strings.ToLower(Produto.Nome) == minusculo{
			ItensEncontrados = append(ItensEncontrados, Produto)
		}
		}
	if len(ItensEncontrados) == 0{
		return nil, errors.New("Nenhum produto encontrado")
	}
	return ItensEncontrados,nil
}
func (r *Estoque) TotalEmEstoque() (totalItens int, TotalEstoque float64){
	for _, Produto := range r.BancoDados{
		totalItens += Produto.Qnt
		TotalEstoque+= float64(Produto.Qnt)*Produto.Preco
	}
	return
}

		
	

		


func main() {
	// Inicializando o Estoque com ID inicial em 1
}