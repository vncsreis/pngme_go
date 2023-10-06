package cmd

import (
	"fmt"
	"os"
	"pngme/png"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:  "remove",
	Args: cobra.ExactArgs(2),
	Run:  remove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func remove(cmd *cobra.Command, args []string) {
	filePath := args[0]
	typeString := args[1]

	pngFile, err := png.FromPath(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %s", filePath, err)
		os.Exit(1)
	}

	err = pngFile.RemoveChunk(typeString)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error removing chunk with type %s: %s",
			filePath,
			err,
		)
		os.Exit(1)
	}

	newFile, err := os.Create(filePath)
	newFile.Write(pngFile.AsBytes())
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Error overwriting file: %s",
			err,
		)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Chunk with type: %s removed", typeString)

	os.Exit(0)
}
