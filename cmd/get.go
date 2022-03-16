/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub repository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		var name = "dr-who.png"

		if len(args) >= 1 && args[0] != "" {
			name = args[0]
		}

		url := "https://github.com/scraly/gophers/raw/main/" + name + ".png"

		fmt.Println("Try to get '" + name + "' Gopher...")

		// Get the data
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			// Create the file
			file, err := os.Create(name + ".png")
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			// Writer the body to file
			_, err = io.Copy(file, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Perfect! Just saved in " + file.Name() + "!")
		} else {
			fmt.Println("Error: " + name + " not exists! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
