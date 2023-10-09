package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pngme",
	Short: "PNGme encodes and decodes strings into PNGs",
	Long: `Encode and decode messages into PNG files, based on chunk types.
    Based on https://picklenerd.github.io/pngme_book/, but built with Go.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error on CLI: %s", err)
		os.Exit(1)
	}
}
