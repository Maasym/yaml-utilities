package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

var input, output, format string

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert YAML to JSON or JSON to YAML",
	Long:  "This command converts a YAML file to JSON or a JSON file to YAML.",
	Run: func(cmd *cobra.Command, args []string) {
		if input == "" || output == "" || format == "" {
			fmt.Println("Error: --input, --output, and --format are required")
			os.Exit(1)
		}

		// Read the input file
		data, err := os.ReadFile(input)
		if err != nil {
			fmt.Printf("Failed to read input file '%s': %v\n", input, err)
			os.Exit(1)
		}

		if format == "json" {
			// Convert YAML to JSON
			var parsedData interface{}
			if err := yaml.Unmarshal(data, &parsedData); err != nil {
				fmt.Printf("Invalid YAML in input file '%s': %v\n", input, err)
				os.Exit(1)
			}
			jsonData, err := json.MarshalIndent(parsedData, "", "  ")
			if err != nil {
				fmt.Printf("Failed to convert to JSON: %v\n", err)
				os.Exit(1)
			}
			if err := os.WriteFile(output, jsonData, 0644); err != nil {
				fmt.Printf("Failed to write output file '%s': %v\n", output, err)
				os.Exit(1)
			}
		} else if format == "yaml" {
			// Convert JSON to YAML
			var parsedData interface{}
			if err := json.Unmarshal(data, &parsedData); err != nil {
				fmt.Printf("Invalid JSON in input file '%s': %v\n", input, err)
				os.Exit(1)
			}
			yamlData, err := yaml.Marshal(parsedData)
			if err != nil {
				fmt.Printf("Failed to convert to YAML: %v\n", err)
				os.Exit(1)
			}
			if err := os.WriteFile(output, yamlData, 0644); err != nil {
				fmt.Printf("Failed to write output file '%s': %v\n", output, err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Error: Invalid format. Use 'json' or 'yaml'.")
			os.Exit(1)
		}

		fmt.Printf("Converted '%s' to '%s' successfully.\n", input, output)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVar(&input, "input", "", "Path to the input file")
	convertCmd.Flags().StringVar(&output, "output", "", "Path to the output file")
	convertCmd.Flags().StringVar(&format, "format", "", "Target format: 'json' or 'yaml'")
}
