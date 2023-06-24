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

// showallCmd represents the showall command
var ShowallCmd = &cobra.Command{
	Use:   "showall",
	Short: "shows all the keys stored in database",

	Run: func(cmd *cobra.Command, args []string) {
		db, err := pebble.Open("vault", &pebble.Options{})
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		iter := db.NewIter(nil)
		defer iter.Close()

		for iter.First(); iter.Valid(); iter.Next() {
			key := iter.Key()
			fmt.Printf("%s\n", key)
		}
	},
}

func init() {

}
