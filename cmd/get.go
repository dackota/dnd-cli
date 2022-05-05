/*
Copyright © 2022 Dackota Johnson dackota.j@gmail.com

*/
package cmd

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all characters",
	Long:  `List all characters currently in the database in JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchCharacters()
	},
}

func fetchCharacters() {
	url := "http://localhost:8080/characters"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyString := string(body)
	log.Print(bodyString)
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
