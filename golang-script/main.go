package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectName string
var environment string
var countries []string

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "An example program that takes a string, a string and a slice as input",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Project Name:", projectName)
		fmt.Println("Environment:", environment)
		fmt.Println("Countries:", countries)
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
