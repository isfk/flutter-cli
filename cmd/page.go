/*
Copyright © 2022 sfk <sfk@live.cn>
*/
package cmd

import (
	"log"

	"github.com/isfk/flutter-cli/generator"
	"github.com/spf13/cobra"
)

// pageCmd represents the page command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "page command create: binding&controller&view&model",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("Please tell me the page name, like 'user'")
			return
		}

		if err := generator.CreatePage("", args[0]); err != nil {
			log.Fatal(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(pageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
