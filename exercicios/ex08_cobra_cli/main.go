package main

// ============================================================
//  EXERCÍCIO 8 — CLI Tool com Cobra
// ============================================================
//
// OBJETIVO:
//   Criar uma ferramenta de linha de comando (CLI) usando o
//   pacote `cobra`, o mesmo usado pelo kubectl, GitHub CLI, etc.
//
// SETUP:
//   go mod init ex08_cobra
//   go get github.com/spf13/cobra
//
// INSTRUÇÕES:
//
//  Crie uma CLI chamada `calcfin` (calculadora financeira) com
//  os seguintes subcomandos:
//
//  ─────────────────────────────────────────────────
//  SUBCOMANDO: calcfin imc
//  ─────────────────────────────────────────────────
//    Flags obrigatórias:
//      --peso   float   (peso em kg)
//      --altura float   (altura em metros)
//    Comportamento:
//      Calcula e imprime o IMC e a classificação
//    Exemplo:
//      go run main.go imc --peso 70 --altura 1.75
//      → IMC: 22.86 | Normal
//
//  ─────────────────────────────────────────────────
//  SUBCOMANDO: calcfin salario
//  ─────────────────────────────────────────────────
//    Flags obrigatórias:
//      --bruto float   (salário bruto)
//    Flag opcional:
//      --inss bool     (descontar INSS, default: true)
//    Comportamento:
//      Calcula o salário líquido com base na tabela INSS:
//        Até R$ 1.412,00   → 7,5%
//        Até R$ 2.666,68   → 9,0%
//        Até R$ 4.000,03   → 12,0%
//        Acima             → 14,0%
//    Exemplo:
//      go run main.go salario --bruto 3500
//      → Bruto: R$ 3500.00 | Desconto: R$ 420.00 | Líquido: R$ 3080.00
//
//  ─────────────────────────────────────────────────
//  SUBCOMANDO: calcfin juros
//  ─────────────────────────────────────────────────
//    Flags obrigatórias:
//      --capital  float   (valor inicial)
//      --taxa     float   (taxa de juros em % ao mês)
//      --meses    int     (número de meses)
//    Comportamento:
//      Calcula juros compostos: M = C * (1 + i)^n
//      Imprime uma tabela mês a mês
//    Exemplo:
//      go run main.go juros --capital 1000 --taxa 1.5 --meses 6
//      Mês 1: R$ 1015.00
//      Mês 2: R$ 1030.23
//      ...
//
//  ESTRUTURA SUGERIDA:
//    main.go        → cria o rootCmd e chama Execute()
//    cmd/imc.go     → define o subcomando imc
//    cmd/salario.go → define o subcomando salario
//    cmd/juros.go   → define o subcomando juros
//
// DICA — Estrutura básica de um comando cobra:
//   var imcCmd = &cobra.Command{
//       Use:   "imc",
//       Short: "Calcula o IMC",
//       Run: func(cmd *cobra.Command, args []string) {
//           peso, _ := cmd.Flags().GetFloat64("peso")
//           // sua lógica aqui
//       },
//   }
//
//   func init() {
//       rootCmd.AddCommand(imcCmd)
//       imcCmd.Flags().Float64("peso", 0, "Peso em kg")
//   }
//
// ============================================================

func main() {
	// Escreva seu código aqui
}
