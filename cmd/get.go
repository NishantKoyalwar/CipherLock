/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"

	"github.com/cockroachdb/pebble"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves stored credentials for a specific website or account from the password manager.",

	Run: func(cmd *cobra.Command, args []string) {
		web, _ := cmd.Flags().GetString("get")
		db, _ := pebble.Open("vault", &pebble.Options{})
		defer db.Close()
		key := []byte(web)
		value, closer, err := db.Get(key)
		if err != nil {
			log.Fatal(err)
		}
		defer closer.Close()
		cipherText := []byte(value)
		Encryptionkey := []byte("passphrasewhichneedstobe32bytess")

		c, err := aes.NewCipher(Encryptionkey)
		if err != nil {
			fmt.Println(err)
		}
		gcm, err := cipher.NewGCM(c)
		if err != nil {
			fmt.Println(err)
		}

		nonceSize := gcm.NonceSize()
		if len(cipherText) < nonceSize {
			fmt.Println(err)
			return
		}

		nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
		plainText, err := gcm.Open(nil, nonce, cipherText, nil)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("====> Password is '%s'", plainText)

	},
}

func init() {
	//rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.
	GetCmd.Flags().StringP("get", "g", "", "to fetch the password")
	GetCmd.MarkFlagRequired("get")

}
