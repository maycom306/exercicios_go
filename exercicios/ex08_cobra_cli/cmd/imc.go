package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//calcular o imc

var peso float64
var altura float64

var imcCmd = &cobra.Command{
	Use:   "imc",
	Short: "Calcula o IMC",
	RunE: func(cmd *cobra.Command, args []string) error {
		//calcular o imc
		if peso <= 0 || altura <= 0 {
			return fmt.Errorf("Peso e/ou altura nao deve ser 0")
		}
		imc := peso / (altura * altura)
		fmt.Printf("IMC: %.2f\n", imc)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(imcCmd)

	imcCmd.Flags().Float64Var(&peso, "peso", 0, "peso em kg")
	imcCmd.Flags().Float64Var(&altura, "altura", 0, "altura em metros")

	imcCmd.MarkFlagRequired("peso")
	imcCmd.MarkFlagRequired("altura")

}
