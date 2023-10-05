package cmd

import (
	"fmt"
	"os"
	"pngme/png"

	"github.com/spf13/cobra"
)

var decode = &cobra.Command{
	Use:  "decode",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		typeString := args[1]

		pngFile, err := png.FromPath(filePath)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"Error opening file %s: %s",
				filePath,
				err,
			)
			os.Exit(1)
		}

		foundChunk := pngFile.GetChunkByType(typeString)
		if foundChunk == nil {
			fmt.Fprintf(
				os.Stdout,
				"Chunk with type %s not found on file",
				typeString,
			)
			os.Exit(1)
		}

		message := foundChunk.DataAsString()

		fmt.Fprintf(os.Stdout, "Meessage:\n%s", message)

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(decode)
}
