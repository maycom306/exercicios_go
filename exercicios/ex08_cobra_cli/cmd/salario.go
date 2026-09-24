package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bruto float64
var inss bool
var irrf bool

var salarioCmd = &cobra.Command{
	Use:   "salario",
	Short: "Calcula o salário líquido com INSS e IRRF",

	RunE: func(cmd *cobra.Command, args []string) error {

		if bruto <= 0 {
			return fmt.Errorf("o salário bruto deve ser um valor positivo")
		}

		// =========================================================
		// INSS
		// =========================================================

		var aliquotaINSS float64
		var deducaoINSS float64
		var descontoINSS float64

		switch {
		case bruto <= 1621.00:
			aliquotaINSS = 0.075
			deducaoINSS = 0

		case bruto <= 2902.84:
			aliquotaINSS = 0.09
			deducaoINSS = 24.32

		case bruto <= 4354.27:
			aliquotaINSS = 0.12
			deducaoINSS = 111.40

		case bruto <= 8475.55:
			aliquotaINSS = 0.14
			deducaoINSS = 198.49

		default:
			aliquotaINSS = 0.14
			descontoINSS = 988.07
		}

		if descontoINSS == 0 {
			descontoINSS = (bruto * aliquotaINSS) - deducaoINSS
		}

		if !inss {
			descontoINSS = 0
		}

		// =========================================================
		// IRRF
		// =========================================================

		var baseIRRF float64
		var aliquotaIRRF float64
		var deducaoIRRF float64
		var impostoIRRF float64

		if irrf {

			// Para este exemplo estamos utilizando o desconto
			// simplificado mensal de R$ 607,20 quando ele for
			// mais vantajoso que as deduções legais.
			//
			// Como estamos considerando apenas o INSS como
			// dedução legal:
			descontoSimplificado := 607.20

			deducaoLegal := descontoINSS

			var deducaoBase float64

			if descontoSimplificado > deducaoLegal {
				deducaoBase = descontoSimplificado
			} else {
				deducaoBase = deducaoLegal
			}

			baseIRRF = bruto - deducaoBase

			if baseIRRF < 0 {
				baseIRRF = 0
			}

			// Tabela progressiva mensal IRRF 2026
			switch {
			case baseIRRF <= 2428.80:
				aliquotaIRRF = 0
				deducaoIRRF = 0

			case baseIRRF <= 2826.65:
				aliquotaIRRF = 0.075
				deducaoIRRF = 182.16

			case baseIRRF <= 3751.05:
				aliquotaIRRF = 0.15
				deducaoIRRF = 394.16

			case baseIRRF <= 4664.68:
				aliquotaIRRF = 0.225
				deducaoIRRF = 675.49

			default:
				aliquotaIRRF = 0.275
				deducaoIRRF = 908.73
			}

			impostoIRRF = (baseIRRF * aliquotaIRRF) - deducaoIRRF

			if impostoIRRF < 0 {
				impostoIRRF = 0
			}

			// =====================================================
			// REDUÇÃO DO IRRF - 2026
			// =====================================================

			var reducao float64

			switch {
			case bruto <= 5000:
				// Para rendimentos de até R$ 5.000,
				// o imposto fica zerado.
				reducao = impostoIRRF

			case bruto <= 7350:
				reducao = 978.62 - (0.133145 * bruto)

				if reducao < 0 {
					reducao = 0
				}

			default:
				reducao = 0
			}

			impostoIRRF -= reducao

			if impostoIRRF < 0 {
				impostoIRRF = 0
			}
		}

		// =========================================================
		// SALÁRIO LÍQUIDO
		// =========================================================

		liquido := bruto - descontoINSS - impostoIRRF

		// =========================================================
		// RESULTADO
		// =========================================================

		fmt.Printf("\n")
		fmt.Printf("+==============================================+\n")
		fmt.Printf("|              CONTRACHEQUE SALARIAL           |\n")
		fmt.Printf("+==============================================+\n")
		fmt.Printf("| Bruto:          R$ %-19.2f |\n", bruto)
		fmt.Printf("| INSS:           R$ %-19.2f |\n", descontoINSS)
		fmt.Printf("| Base IRRF:      R$ %-19.2f |\n", baseIRRF)
		fmt.Printf("| IRRF:           R$ %-19.2f |\n", impostoIRRF)
		fmt.Printf("--------------------------------------------\n")
		fmt.Printf("| Líquido:        R$ %-19.2f |\n", liquido)
		fmt.Printf("+==============================================+\n")

		return nil
	},
}

func init() {

	rootCmd.AddCommand(salarioCmd)

	salarioCmd.Flags().Float64Var(
		&bruto,
		"bruto",
		0,
		"salário bruto",
	)

	salarioCmd.Flags().BoolVar(
		&inss,
		"inss",
		true,
		"descontar INSS",
	)

	salarioCmd.Flags().BoolVar(
		&irrf,
		"irrf",
		true,
		"descontar IRRF",
	)

	salarioCmd.MarkFlagRequired("bruto")
}
