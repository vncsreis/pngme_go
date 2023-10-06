package cmd

import (
	"fmt"
	"os"
	"pngme/png"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:  "decode",
	Args: cobra.ExactArgs(2),
	Run:  decode,
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

func decode(cmd *cobra.Command, args []string) {
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
			"Error: Chunk with type %s not found on file\n",
			typeString,
		)
		os.Exit(1)
	}

	message := foundChunk.DataAsString()

	fmt.Fprintf(os.Stdout, "Meessage:\n%s\n", message)

	os.Exit(0)
}
