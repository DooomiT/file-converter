package cmd

import (
	"os"
	"strings"

	"github.com/dooomit/file-converter/pkg/convert"
	"github.com/spf13/cobra"
)

func CArray() *cobra.Command {
	command := &cobra.Command{
		Use:   "c-array",
		Short: "Converts a file to a C array",
		Long:  `Converts a file to a C array. The array can be used in C/C++ programs. You can specify the format of the array with the --format flag. The default format is "hex". The array is written to the output file. If no output file is given, the array is written to <input>.c.`,
		Run: func(cmd *cobra.Command, args []string) {
			inputFile := getRequiredFlag(cmd, "input")
			outputFile := getOptionalFlag(cmd, "output", inputFile+".c")
			format := getOptionalFlag(cmd, "format", "hex")
			content, err := os.ReadFile(inputFile)
			if err != nil {
				panic(err)
			}
			array := convert.ArrayImpl{
				DataType:     convert.ConversionType(format),
				VariableType: "const unsigned char",
				Name:         strings.Replace(inputFile, ".", "_", -1),
				Data:         content,
			}
			convertedArray, err := array.ConvertToArray()
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(outputFile, []byte(convertedArray), 0644)
			if err != nil {
				panic(err)
			}
			println("Wrote array to file:", outputFile)
		},
	}
	command.Flags().StringP("input", "i", "", "Input file")
	command.Flags().StringP("output", "o", "", "Output file, default is <input>.c")
	command.Flags().StringP("format", "f", "hex", "Format of the array, default is hex")
	command.MarkFlagRequired("input")
	command.MarkFlagFilename("input")
	command.MarkFlagFilename("output")
	return command
}

func getRequiredFlag(cmd *cobra.Command, name string) string {
	value := cmd.Flag(name).Value.String()
	if value == "" {
		println("Missing required flag:", name)
		os.Exit(1)
	}
	return value
}

func getOptionalFlag(cmd *cobra.Command, name string, defaultValue string) string {
	value := cmd.Flag(name).Value.String()
	if value == "" {
		return defaultValue
	}
	return value
}
