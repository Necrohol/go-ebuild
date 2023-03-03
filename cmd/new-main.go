package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"regexp"
	"strings"
	"path/filepath"
	"github.com/pelletier/go-toml"
	"go-dep-parser/blob/main/pkg/golang/mod/parse.go"
	"aquasecurity/go-dep-parser/blob/main/pkg/golang/sum/parse.go" 
)

var rootCmd = &cobra.Command{
	Use:   "go-ebuild",
	Short: "A tool for managing Gentoo ebuilds written in Go",
	Long:  `A CLI tool for managing Gentoo ebuilds written in Go, powered by Cobra and Viper.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to go-ebuild!")
	},
}

var egosumCmd = &cobra.Command{
	Use:   "egosum",
	Short: "Calculate the checksum of dependencies specified in go.mod",
	Long:  `A command that calculates the checksum of dependencies specified in go.mod using the 'get-ego-vendor' command/stores to:EGO_SUM environment variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the "get-ego-vendor" command is available.
		_, err := exec.LookPath("get-ego-vendor")
		if err != nil {
			// Command not found, ask user to install it.
			fmt.Println("get-ego-vendor command not found. Do you want to install it? (Y/N)")
			var answer string
			fmt.Scanln(&answer)
			if strings.ToUpper(answer) == "Y" {
				// Install the command.
				cmd := exec.Command("emerge", "-bavgk", "get-ego-vendor")
				err := cmd.Run()
				if err != nil {
					fmt.Println("Failed to install get-ego-vendor:", err)
					return
				}
				fmt.Println("get-ego-vendor installed successfully.")
			} else {
				fmt.Println("Aborted.")
				return
			}
		}

		// Execute the "get-ego-vendor" command.
		cmd := exec.Command("get-ego-vendor")
		// Create a temporary file to store the output of the command.
		f, err := os.CreateTemp("", "egosum")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(f.Name())

		// Write the output of the command to the temporary file.
		cmd.Stdout = f
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		// Read the contents of the temporary file.
		content, err := os.ReadFile(f.Name())
		if err != nil {
			log.Fatal(err)
		}

		// Store the output in the EGO_SUM environment variable.
		os.Setenv("EGO_SUM", string(content))

		// Print the output to the console.
		fmt.Println(string(content))
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(egosumCmd)
}

func initConfig() {
	viper.SetConfigFile(".go-ebuild.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
``



flag.Var(&go-ebuild, "--git", "read metadata by cloning a git repo to ./tmp or users ./tmp ")

// append to end of package root Program//
// add a gracefull exit 
// I would add a random Quote n exit but for now keep it simple 

var rootCmd = &cobra.Command{
	Use:   "go-ebild Exit",
	Short: "Quit/Exit This utility",
	Long:  `A CLI tool for managing Gentoo ebuilds written in Go, powered by Cobra and Viper.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Have A Nice Day!!, Hopefully You have Enjoied using go-ebuild!")
		 os.Exit(1)
	},
}
