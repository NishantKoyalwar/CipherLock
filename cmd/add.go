/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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
		if err := Db.Set(key, Value, pebble.Sync); err != nil {
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
