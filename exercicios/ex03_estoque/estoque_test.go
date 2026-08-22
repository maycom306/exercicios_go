package main

import (
	"strings"
	"testing"
)

// TestAdcionarProduto verifica se os produtos estão sendo adicionados de forma correta
// e se o auto-incremento do ID único e o registro no mapa funcionam perfeitamente.
func TestAdcionarProduto(t *testing.T) {
	estoque := Estoque{
		BancoDados: make(map[int]Produto),
		ItemID:     1,
	}

	// Adicionando primeiro produto
	p1 := estoque.AdcionarProduto("Teclado", 150.00, 10)
	if p1.ID != 1 {
		t.Errorf("Esperado ID 1 para o primeiro produto, obtido %d", p1.ID)
	}
	if estoque.ItemID != 2 {
		t.Errorf("Esperado ItemID atualizado para 2, obtido %d", estoque.ItemID)
	}

	// Adicionando segundo produto
	p2 := estoque.AdcionarProduto("Mouse", 80.00, 5)
	if p2.ID != 2 {
		t.Errorf("Esperado ID 2 para o segundo produto, obtido %d", p2.ID)
	}

	// Verifica se ambos estão no mapa
	if len(estoque.BancoDados) != 2 {
		t.Errorf("Esperado tamanho do mapa igual a 2, obtido %d", len(estoque.BancoDados))
	}
}

// TestRemoverEstoque garante o controle correto de saídas físicas, 
// tratando tanto os casos de sucesso quanto as validações de erro obrigatórias.
func TestRemoverEstoque(t *testing.T) {
	estoque := Estoque{
		BancoDados: make(map[int]Produto),
		ItemID:     1,
	}

	// Adicionando produto para teste de remoção (ID será 1)
	estoque.AdcionarProduto("Teclado", 150.00, 10)

	// Cenário 1: Remoção de estoque com sucesso
	err := estoque.RemoverEstoque(1, 4)
	if err != nil {
		t.Fatalf("Erro inesperado em remoção válida: %v", err)
	}
	if prod := estoque.BancoDados[1]; prod.Qnt != 6 {
		t.Errorf("Esperado estoque de 6 itens restantes, obtido %d", prod.Qnt)
	}

	// Cenário 2: Erro de estoque insuficiente
	err = estoque.RemoverEstoque(1, 15)
	if err == nil {
		t.Error("Esperava erro por estoque insuficiente, mas a remoção foi permitida")
	} else if !strings.Contains(err.Error(), "Estoque Insuficiente") {
		t.Errorf("Mensagem de erro inesperada para estoque insuficiente: %v", err)
	}

	// Cenário 3: Erro de produto inexistente
	err = estoque.RemoverEstoque(999, 1)
	if err == nil {
		t.Error("Esperava erro de item não encontrado, mas nenhum erro foi retornado")
	} else if !strings.Contains(err.Error(), "item Não encontrado") {
		t.Errorf("Mensagem de erro inesperada para item não encontrado: %v", err)
	}
}

// TestBuscarPorNome verifica a busca case-insensitive e o retorno de múltiplos itens,
// além de validar o erro quando o produto não existe.
func TestBuscarPorNome(t *testing.T) {
	estoque := Estoque{
		BancoDados: make(map[int]Produto),
		ItemID:     1,
	}

	estoque.AdcionarProduto("Teclado Mecanico", 150.00, 10)
	estoque.AdcionarProduto("teclado de membrana", 80.00, 5)
	estoque.AdcionarProduto("Mouse", 50.00, 3)

	// Cenário 1: Busca exata ignorando maiúsculas/minúsculas
	resultados, err := estoque.BuscarPorNome("TECLADO MECANICO")
	if err != nil {
		t.Fatalf("Busca por 'TECLADO MECANICO' falhou inesperadamente: %v", err)
	}
	if len(resultados) != 1 || resultados[0].ID != 1 {
		t.Errorf("Esperado encontrar 1 produto com ID 1, obtido: %v", resultados)
	}

	// Cenário 2: Busca por nome que retorna múltiplos itens
	// Nota: Se você usar busca exata de string == strings.ToLower(Nome), o teste abaixo 
	// garante o funcionamento exato do termo. Se o exercício exigir strings.Contains,
	// você pode alterar o método para suportar buscas parciais.
	resultadosMin, err := estoque.BuscarPorNome("teclado de membrana")
	if err != nil {
		t.Fatalf("Busca por 'teclado de membrana' falhou: %v", err)
	}
	if len(resultadosMin) != 1 {
		t.Errorf("Esperado encontrar 1 produto, obtido %d", len(resultadosMin))
	}

	// Cenário 3: Busca por nome inexistente (deve retornar erro)
	_, err = estoque.BuscarPorNome("Monitor")
	if err == nil {
		t.Error("Esperava erro de 'Nenhum produto encontrado', mas a busca retornou sucesso")
	} else if !strings.Contains(err.Error(), "Nenhum produto encontrado") {
		t.Errorf("Mensagem de erro de busca inesperada: %v", err)
	}
}

// TestTotalEmEstoque garante o cálculo matemático correto de todo o estoque (Unidades e Valor Financeiro)
func TestTotalEmEstoque(t *testing.T) {
	estoque := Estoque{
		BancoDados: make(map[int]Produto),
		ItemID:     1,
	}

	// Adicionando produtos controlados
	estoque.AdcionarProduto("Teclado", 150.00, 10) // Valor: 1500.00
	estoque.AdcionarProduto("Mouse", 80.00, 5)     // Valor: 400.00

	totalItens, totalEstoque := estoque.TotalEmEstoque()

	// 1. Validando quantidade física total
	if totalItens != 15 {
		t.Errorf("Esperado total de 15 itens físicos, obtido %d", totalItens)
	}

	// 2. Validando valor financeiro acumulado (Preço * Quantidade)
	valorEsperado := 1900.00
	if totalEstoque != valorEsperado {
		t.Errorf("Esperado valor financeiro acumulado de R$ %.2f, obtido R$ %.2f", valorEsperado, totalEstoque)
	}
}
