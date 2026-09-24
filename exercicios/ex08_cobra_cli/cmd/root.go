package cmd

import ("github.com/spf13/cobra")

var rootCmd = &cobra.Command{
	Use:   "Calculadora",
	Short: "Calculadora de IMC, Salário e Juros",
}

func Execute() error {
	return rootCmd.Execute()
}
	