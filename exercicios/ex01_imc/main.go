package main

// ============================================================
//  EXERCÍCIO 1 — Calculadora de IMC com Structs e Métodos
// ============================================================
//
// OBJETIVO:
//   Praticar a criação de structs, métodos e switch em Go.
//
// INSTRUÇÕES:
//
//  1. Crie uma struct chamada `Pessoa` com os campos:
//       - Nome   string
//       - Peso   float64   (em kg)
//       - Altura float64   (em metros)
//
//  2. Implemente um MÉTODO na struct chamado `IMC()` que:
//       - Calcula e retorna o IMC usando: Peso / (Altura * Altura)
//
//  3. Implemente um MÉTODO chamado `Classificacao()` que:
//       - Retorna uma string conforme a tabela abaixo:
//         IMC < 18.5         → "Abaixo do peso"
//         18.5 <= IMC < 25.0 → "Normal"
//         25.0 <= IMC < 30.0 → "Sobrepeso"
//         IMC >= 30.0        → "Obesidade"
//
//  4. No main(), crie uma slice com pelo menos 5 pessoas
//     e itere sobre ela imprimindo: nome, IMC e classificação.
//
// SAÍDA ESPERADA (exemplo):
//   Ana     | IMC: 22.50 | Normal
//   Carlos  | IMC: 30.12 | Obesidade
//   ...
//
// DICA:
//   Métodos em Go são declarados assim:
//     func (p Pessoa) NomeDoMetodo() TipoRetorno { ... }
//
// ============================================================

func main() {
	// Escreva seu código aqui
}
