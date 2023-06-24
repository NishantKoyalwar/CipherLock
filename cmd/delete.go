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

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes stored credentials from the password manager for a specified website or account.",

	Run: func(cmd *cobra.Command, args []string) {
		del, _ := cmd.Flags().GetString("delete")

		db, _ := pebble.Open("vault", &pebble.Options{})
		defer db.Close()
		if err := db.Delete([]byte(del), &pebble.WriteOptions{}); err != nil {
			log.Fatal(err)
		}

		fmt.Println("credentials deleted successfully")

	},
}

func init() {

	// Here you will define your flags and configuration settings.
	DeleteCmd.Flags().StringP("delete", "d", "", "deletes the credentials")

	DeleteCmd.MarkFlagRequired("delete")

}
