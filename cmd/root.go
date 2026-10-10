package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "unco",
	Short: "unco is a cli unit conversions calculator for engineering applications",
	Long: "unco is a cli tool for advanced unit conversions calculator for engineering applications, such as powers of units and conversions between SI and USCS",
	Run: func (cmd *cobra.Command, args []string) {

	},
}


func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing unco '%s'\n", err)
        os.Exit(1)
    }
}