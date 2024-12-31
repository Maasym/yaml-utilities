package cmd

import (
	"fmt"
	"os"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var file1, file2 string

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Compare two YAML files",
	Long:  "This command compares two YAML files and highlights their differences.",
	Run: func(cmd *cobra.Command, args []string) {
		if file1 == "" || file2 == "" {
			fmt.Println("Error: --file1 and --file2 are required")
			os.Exit(1)
		}

		// Read and parse the first file
		data1, err := os.ReadFile(file1)
		if err != nil {
			fmt.Printf("Failed to read file1 '%s': %v\n", file1, err)
			os.Exit(1)
		}
		var parsed1 interface{}
		if err := yaml.Unmarshal(data1, &parsed1); err != nil {
			fmt.Printf("Invalid YAML in file1 '%s': %v\n", file1, err)
			os.Exit(1)
		}

		// Read and parse the second file
		data2, err := os.ReadFile(file2)
		if err != nil {
			fmt.Printf("Failed to read file2 '%s': %v\n", file2, err)
			os.Exit(1)
		}
		var parsed2 interface{}
		if err := yaml.Unmarshal(data2, &parsed2); err != nil {
			fmt.Printf("Invalid YAML in file2 '%s': %v\n", file2, err)
			os.Exit(1)
		}

		// Compare the files
		diff := cmp.Diff(parsed1, parsed2)
		if diff == "" {
			fmt.Println("The two files are identical.")
		} else {
			fmt.Println("Differences:")
			fmt.Println(diff)
		}
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	diffCmd.Flags().StringVar(&file1, "file1", "", "Path to the first YAML file")
	diffCmd.Flags().StringVar(&file2, "file2", "", "Path to the second YAML file")
}
