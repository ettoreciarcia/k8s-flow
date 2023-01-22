package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var projectName string
var environment string
var countries []string

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "An example program that takes a string, a string and a slice as input",
	Run: func(cmd *cobra.Command, args []string) {
		if projectName != "app1" && projectName != "app2" {
			fmt.Println(errors.New("Error: projectName must be either app1 or app2"))
			return
		}
		if environment != "production" && environment != "staging" {
			fmt.Println(errors.New("Error: environment must be either production or staging"))
			return
		}
		for _, country := range countries {
			if len(country) > 3 {
				fmt.Println(errors.New("Error: country must not be longer than 3 characters"))
				return
			}
		}
		fmt.Println("Project Name:", projectName)
		fmt.Println("Environment:", environment)
		fmt.Println("Countries:", countries)

		var app interface{}
		if projectName == "app1" {
			app = struct {
				FirstApp struct {
					Create    bool     `yaml:"create"`
					AppName   string   `yaml:"appName"`
					Env       string   `yaml:"env"`
					Countries []string `yaml:"countries"`
				} `yaml:"firstApp"`
			}{
				FirstApp: struct {
					Create    bool     `yaml:"create"`
					AppName   string   `yaml:"appName"`
					Env       string   `yaml:"env"`
					Countries []string `yaml:"countries"`
				}{
					Create:    true,
					AppName:   projectName,
					Env:       environment,
					Countries: countries,
				},
			}
		} else {
			app = struct {
				SecondApp struct {
					Create    bool     `yaml:"create"`
					AppName   string   `yaml:"appName"`
					Env       string   `yaml:"env"`
					Countries []string `yaml:"countries"`
				} `yaml:"secondApp"`
			}{
				SecondApp: struct {
					Create    bool     `yaml:"create"`
					AppName   string   `yaml:"appName"`
					Env       string   `yaml:"env"`
					Countries []string `yaml:"countries"`
				}{
					Create:    true,
					AppName:   projectName,
					Env:       environment,
					Countries: countries,
				},
			}
		}

		file, _ := yaml.Marshal(app)
		_ = ioutil.WriteFile(projectName+".yml", file, 0644)

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
