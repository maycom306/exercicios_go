package cmd

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// SUBCOMANDO: calcfin juros
//
// Calcula juros compostos:
// M = C * (1 + i)^n
//
// Exemplo:
//
//	go run main.go juros
//
// Digite o capital: R$ 1000
// Digite a taxa (% ao mês): 1,5
// Digite a quantidade de meses: 6
//
// Mês 1: R$ 1015.00
// Mês 2: R$ 1030.23
// ...

var jurosCmd = &cobra.Command{
	Use:   "juros",
	Short: "Calcula juros compostos mês a mês",

	RunE: func(cmd *cobra.Command, args []string) error {

		// =========================
		// CAPITAL
		// =========================

		fmt.Print("Digite o capital: R$ ")

		var entrada string
		fmt.Scan(&entrada)

		valor, err := converterNumero(entrada)

		if err != nil {
			return fmt.Errorf("capital inválido: %s", entrada)
		}

		// =========================
		// TAXA
		// =========================

		fmt.Print("Digite a taxa (% ao mês): ")

		fmt.Scan(&entrada)

		taxa, err := converterNumero(entrada)

		if err != nil {
			return fmt.Errorf("taxa inválida: %s", entrada)
		}

		// =========================
		// MESES
		// =========================

		fmt.Print("Digite a quantidade de meses: ")

		fmt.Scan(&entrada)

		meses, err := strconv.Atoi(entrada)

		if err != nil {
			return fmt.Errorf("quantidade de meses inválida: %s", entrada)
		}

		// =========================
		// VALIDAÇÃO
		// =========================

		if valor <= 0 {
			return fmt.Errorf("o capital deve ser maior que zero")
		}

		if taxa <= 0 {
			return fmt.Errorf("a taxa deve ser maior que zero")
		}

		if meses <= 0 {
			return fmt.Errorf("a quantidade de meses deve ser maior que zero")
		}

		// =========================
		// CÁLCULO
		// =========================

		CalcularJuros := func(
			valor float64,
			taxa float64,
			meses int,
		) float64 {
			return valor * math.Pow(
				1+taxa/100,
				float64(meses),
			)
		}

		// =========================
		// RESULTADO
		// =========================

		fmt.Println()
		fmt.Println("========== RESULTADO ==========")

		for mes := 1; mes <= meses; mes++ {

			total := CalcularJuros(
				valor,
				taxa,
				mes,
			)

			fmt.Printf(
				"Mês %d: R$ %.2f\n",
				mes,
				total,
			)
		}

		return nil
	},
}

// converterNumero aceita números nos formatos:
//
// 1.5
// 1,5
// 1500.50
// 1500,50
// 1.500,50
// 1,500.50
func converterNumero(entrada string) (float64, error) {

	entrada = strings.TrimSpace(entrada)

	temVirgula := strings.Contains(entrada, ",")
	temPonto := strings.Contains(entrada, ".")

	switch {

	// Exemplo:
	// 1.500,50
	//
	// A vírgula está depois do ponto,
	// então consideramos formato brasileiro.
	case temVirgula && temPonto &&
		strings.LastIndex(entrada, ",") >
			strings.LastIndex(entrada, "."):

		entrada = strings.ReplaceAll(entrada, ".", "")
		entrada = strings.ReplaceAll(entrada, ",", ".")

	// Exemplo:
	// 1,500.50
	//
	// O ponto está depois da vírgula,
	// então consideramos formato americano.
	case temVirgula && temPonto:

		entrada = strings.ReplaceAll(entrada, ",", "")

	// Exemplo:
	// 1,5
	// 1500,50
	case temVirgula:

		entrada = strings.ReplaceAll(entrada, ",", ".")

	// Exemplo:
	// 1.5
	// 1500.50
	//
	// Já está no formato aceito pelo Go.
	}

	return strconv.ParseFloat(entrada, 64)
}

func init() {
	rootCmd.AddCommand(jurosCmd)
}