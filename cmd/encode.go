package cmd

import (
	"fmt"
	"os"
	"pngme/chunks"
	"pngme/png"

	"github.com/spf13/cobra"
)

// TODO: update file contents duh
var encode = &cobra.Command{
	Use:  "encode",
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		typeString := args[1]
		message := args[2]

		pngFile, err := png.FromPath(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %s", filePath, err)
			os.Exit(1)
		}

		newChunk, err := chunks.FromStrings(typeString, message)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating new chunk: %s", err)
		}

		pngFile.AppendChunk(*newChunk)

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(encode)
}
