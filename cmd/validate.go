package cmd

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

var (
	path         string
	excludePaths []string
	outputFormat string
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate YAML files",
	Long:  "This command validates YAML files or all YAML files in a directory to ensure they are properly formatted.",
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {
			fmt.Println("Error: --path is required")
			os.Exit(1)
		}

		// Check if the path is a file or directory
		info, err := os.Stat(path)
		if err != nil {
			fmt.Printf("Error: Could not access path '%s': %v\n", path, err)
			os.Exit(1)
		}

		if info.IsDir() {
			validateDirectory(path)
		} else {
			validateFile(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVarP(&path, "path", "p", "", "Path to a YAML file or directory")
	validateCmd.Flags().StringSliceVar(&excludePaths, "exclude", []string{}, "Files or directories to exclude from validation")
	validateCmd.Flags().StringVar(&outputFormat, "output", "text", "Output format: text, json, yaml")
}

// validateFile validates a single YAML file
func validateFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file '%s': %v\n", filePath, err)
		printSummary(1, 0, 1, []string{fmt.Sprintf("Failed to read file '%s': %v", filePath, err)})
		return
	}

	var parsedData interface{}
	err = yaml.Unmarshal(data, &parsedData)
	if err != nil {
		fmt.Printf("Invalid YAML in file '%s': %v\n", filePath, err)
		printSummary(1, 0, 1, []string{fmt.Sprintf("Invalid YAML in file '%s': %v", filePath, err)})
		return
	}

	fmt.Printf("The file '%s' is a valid YAML file.\n", filePath)
	printSummary(1, 1, 0, nil) // One file validated, no errors
}

// validateDirectory validates all YAML files in a directory
func validateDirectory(dirPath string) {
	var wg sync.WaitGroup
	fileChan := make(chan string, 100) // Buffered channel to avoid blocking
	errorChan := make(chan string, 100)
	var totalFiles, validFiles, invalidFiles int
	mu := sync.Mutex{}

	// Worker Goroutine
	go func() {
		for filePath := range fileChan {
			err := validateFileWithErrorHandling(filePath)
			mu.Lock()
			if err != nil {
				invalidFiles++
				errorChan <- fmt.Sprintf("Invalid YAML in file '%s': %v", filePath, err)
			} else {
				validFiles++
			}
			mu.Unlock()
			wg.Done() // Decrement WaitGroup when the file is processed
		}
		close(errorChan) // Close errorChan after processing all files
	}()

	// Walk the directory and send files to the channel
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Skip excluded paths
		for _, exclude := range excludePaths {
			if filepath.HasPrefix(path, exclude) {
				return nil
			}
		}

		// Add YAML files to processing
		if !info.IsDir() && (filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml") {
			mu.Lock()
			totalFiles++
			mu.Unlock()
			wg.Add(1) // Increment WaitGroup counter for each file
			fileChan <- path
		}
		return nil
	})

	// Close fileChan after all files are sent
	if err != nil {
		fmt.Printf("Error walking directory '%s': %v\n", dirPath, err)
	}
	close(fileChan)

	// Wait for all workers to finish
	wg.Wait()

	// Collect errors
	errorList := []string{}
	for err := range errorChan {
		errorList = append(errorList, err)
	}

	// Print the summary
	printSummary(totalFiles, validFiles, invalidFiles, errorList)
}

func validateFileWithErrorHandling(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var parsedData interface{}
	err = yaml.Unmarshal(data, &parsedData)
	if err != nil {
		return fmt.Errorf("invalid YAML: %w", err)
	}

	return nil
}

// printSummary outputs the validation summary
func printSummary(total, valid, invalid int, errors []string) {
	switch outputFormat {
	case "json":
		printSummaryAsJSON(total, valid, invalid, errors)
	case "yaml":
		printSummaryAsYAML(total, valid, invalid, errors)
	default:
		printSummaryAsText(total, valid, invalid, errors)
	}
}

func printSummaryAsText(total, valid, invalid int, errors []string) {
	fmt.Printf("\nValidation Summary:\n")
	fmt.Printf("Total files checked: %d\n", total)
	fmt.Printf("Valid files: %d\n", valid)
	fmt.Printf("Invalid files: %d\n", invalid)
	if len(errors) > 0 {
		fmt.Println("\nErrors:")
		for _, e := range errors {
			fmt.Println(e)
		}
	}
}

func printSummaryAsJSON(total, valid, invalid int, errors []string) {
	result := map[string]interface{}{
		"total":   total,
		"valid":   valid,
		"invalid": invalid,
		"errors":  errors,
	}
	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func printSummaryAsYAML(total, valid, invalid int, errors []string) {
	result := map[string]interface{}{
		"total":   total,
		"valid":   valid,
		"invalid": invalid,
		"errors":  errors,
	}
	data, _ := yaml.Marshal(result)
	fmt.Println(string(data))
}
