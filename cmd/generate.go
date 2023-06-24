/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a strong and secure password with specified length and complexity",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("...")
		url := "https://password-generator18.p.rapidapi.com/generate-password"

		payload := strings.NewReader("{\n    \"length\": 15,\n    \"numbers\": true,\n    \"symbols\": false,\n    \"lowercase\": true,\n    \"uppercase\": true,\n    \"excludeSimilarCharacters\": false,\n    \"exclude\": \"1\",\n    \"strict\": false\n}")

		req, _ := http.NewRequest("POST", url, payload)
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		URL := os.Getenv("URL")
		KEY := os.Getenv("API_KEY")

		req.Header.Add("content-type", "application/json")
		req.Header.Add("X-RapidAPI-Key", KEY)
		req.Header.Add("X-RapidAPI-Host", URL)

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		//fmt.Println(res)
		fmt.Println(string(body))
	},
}

func init() {

}
