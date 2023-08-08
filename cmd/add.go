/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"github.com/cockroachdb/pebble"

	"github.com/spf13/cobra"
)

type Value struct {
	username string
	password string
}

var AddCmd = &cobra.Command{

	Use:   "add",
	Short: "Adds website credentials to the password manager, providing the website name and password.",

	Run: func(cmd *cobra.Command, args []string) {

		web, _ := cmd.Flags().GetString("Website")
		pas, _ := cmd.Flags().GetString("Password")

		//fmt.Println(web, pas)

		Db, err := pebble.Open("vault", &pebble.Options{})
		if err != nil {
			log.Fatal(err)
		}
		key := []byte(web)
		Value := []byte(pas)
		Encryptionkey := []byte("passphrasewhichneedstobe32bytess")

		//generate a new aes cipher using our encryptionkey
		c, err := aes.NewCipher(Encryptionkey)
		if err != nil {
			fmt.Println(err)
		}
		//gcm or galios/counter mode,is a mode of operation for symm key cryptography block ciphers

		gcm, err := cipher.NewGCM(c)
		if err != nil {
			fmt.Println(err)
		}
		//create a new byte array the size of the nonce which must be passed to seal
		nonce := make([]byte, gcm.NonceSize())

		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			fmt.Println(err)
		}

		if err := Db.Set(key, gcm.Seal(nonce, nonce, Value, nil), pebble.Sync); err != nil {
			log.Fatal(err)
		}

		if err := Db.Close(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("====>credentials saved successfully")

	},
}

func init() {

	AddCmd.PersistentFlags().String("foo", "", "A help for foo")
	AddCmd.Flags().StringP("Website", "w", "", "Website Name")
	AddCmd.Flags().StringP("Password", "p", "", "password for the website")
	AddCmd.MarkFlagRequired("website")
	AddCmd.MarkFlagRequired("Password")

}
