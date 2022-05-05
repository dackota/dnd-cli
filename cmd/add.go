/*
Copyright © 2022 Dackota Johnson dackota.j@gmail.com

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Imports a character from DND Beyond. Requires the character ID.",
	Long:  `Import a character from DND Beyond using the character ID. Can add multiple characters at once.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, characterId := range args {
			characterJson := fetchCharacterJson(characterId)
			postCharacter(characterJson)
		}
	},
}

func fetchCharacterJson(characterId string) []byte {
	url := "https://character-service.dndbeyond.com/character/v3/character/" + characterId
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Couldn't find character data\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func postCharacter(characterJson []byte) {
	url := "http://localhost:8080/characters"

	// Unmarshal
	var character Character
	err := json.Unmarshal(characterJson, &character)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Got character data for: ", character.Data.Name)

	output, err := json.Marshal(character)
	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(output))
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
