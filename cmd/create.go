/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/isfk/flutter-cli/generator"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a flutter app with get",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("Please tell me the app name, like 'myapp'")
			return
		}
		log.Println(cmd.Flag("org").Value)

		log.Println("生成路径:", fmt.Sprintf("./%v", args[0]))

		// 使用 flutter 命令创建
		appCmd := exec.Command("flutter", "create", fmt.Sprintf("--org"), cmd.Flag("org").Value.String(), args[0])
		if err := appCmd.Run(); err != nil {
			log.Fatal(err)
			return
		}

		if err := generator.CreateFiles(args[0]); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.PersistentFlags().String("org", "", `The organization responsible for your new Flutter project, in reverse domain name notation. This string is used in Java package names and as prefix in the iOS bundle identifier.
	(defaults to "com.example")`)
}
