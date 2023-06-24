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

		fmt.Printf("====> Password is '%s'", value)

	},
}

func init() {
	//rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.
	GetCmd.Flags().StringP("get", "g", "", "to fetch the password")
	GetCmd.MarkFlagRequired("get")

}
