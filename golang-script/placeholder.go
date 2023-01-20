package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

var projectName string
var environment string
var countries []string

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "An example program that takes a string, a string and a slice as input",
	Run: func(cmd *cobra.Command, args []string) {
		// Read the contents of the existing file
		data, err := ioutil.ReadFile("template.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Replace placeholders in the file contents with the input values
		fileContent := strings.Replace(string(data), "{{projectName}}", projectName, -1)
		fileContent = strings.Replace(fileContent, "{{environment}}", environment, -1)
		fileContent = strings.Replace(fileContent, "{{countries}}", strings.Join(countries, ", "), -1)

		// Write the new contents to a new file
		err = ioutil.WriteFile("output.txt", []byte(fileContent), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Println("output.txt file created with provided input")
	},
}

func init() {
	rootCmd.Flags().StringVarP(&projectName, "projectName", "p", "", "A string input")
	rootCmd.MarkFlagRequired("projectName")
	rootCmd.Flags().StringVarP(&environment, "environment", "e", "", "A string input")
	rootCmd.MarkFlagRequired("environment")
	rootCmd.Flags().StringSliceVarP(&countries, "countries", "c", []string{}, "A slice of countries input")
	rootCmd.MarkFlagRequired("countries")
}

func main() {
	rootCmd.Execute()
}
