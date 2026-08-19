package main

import (
	"fmt"
	"errors"
)

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
type pessoa struct{
	nome string
	kg float32
	alt float32
}

func (p pessoa) cal()(float32, error){
		if p.alt <= 0{
			return 0, errors.New("Altura tem que superior a 0")
		}
		if p.kg <= 0{
			return 0, errors.New("Peso tem que ser mais que 0)")
		}
		imc := p.kg/(p.alt*p.alt)
		return imc, nil
}
func (p pessoa)class(imc float32)string{
	switch{
	case imc <18.5:
		return "abaixo do peso"
	case imc >= 18.5 && imc < 25.0:
		return "Peso normal"
	case imc >= 25.0 && imc < 30.0:
		return "Sobrepeso"
	case imc >= 30.0:
		return "Obesidade"
	default:
		return "Dados invalidos"
	}
}




func main(){
	// Slice contendo 5 pessoas válidas e 1 caso de teste para gerar erro
	pessoas := []pessoa{
		{nome: "Ana", kg: 52.0, alt: 1.63},
		{nome: "Bruno", kg: 80.0, alt: 1.75},
		{nome: "Carla", kg: 68.0, alt: 1.58},
		{nome: "Diego", kg: 95.0, alt: 1.82},
		{nome: "Elena", kg: 58.0, alt: 1.70},
		{nome: "Teste Inválido", kg: 70.0, alt: 0.0}, // <- Caso de teste com erro (divisão por zero)
	}

	for _, p := range pessoas {
		imc, err := p.cal()

		if err != nil {
			fmt.Printf("❌ [%s] Erro: %s\n", p.nome, err)
			continue
		}

		classificacao := p.class(imc)
		fmt.Printf("✅ [%s] | IMC: %.2f | %s\n", p.nome, imc, classificacao)
	}
}
