package cmd

import (
	"fmt"
	"os"
	"pngme/chunks"
	"pngme/png"

	"github.com/spf13/cobra"
)

var output string

var encodeCmd = &cobra.Command{
	Use:  "encode",
	Args: cobra.ExactArgs(3),
	Run:  encode,
}

func init() {
	encodeCmd.Flags().StringVarP(&output, "output", "o", "", "Output file (if empty program will overwrite file)")
	rootCmd.AddCommand(encodeCmd)
}

func encode(cmd *cobra.Command, args []string) {
	filePath := args[0]
	typeString := args[1]
	message := args[2]

	pngFile, err := png.FromPath(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %s\n", filePath, err)
		os.Exit(1)
	}

	newChunk, err := chunks.FromStrings(typeString, message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating new chunk: %s\n", err)
		os.Exit(1)
	}

	pngFile.AppendChunk(*newChunk)

	if output != "" {
		filePath = output
	}

	newFile, err := os.Create(filePath)
	newFile.Write(pngFile.AsBytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error overwriting file: %s\n", err)
	}

	fmt.Println("Write successful")

	os.Exit(0)
}
