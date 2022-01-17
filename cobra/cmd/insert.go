/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name string = ""
	age  int    = 0
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(name, age)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringVarP(&name, "name", "n", "", "name to insert")
	insertCmd.MarkFlagRequired("name")

	insertCmd.Flags().IntVarP(&age, "age", "a", 0, "age to insert")
	insertCmd.MarkFlagRequired("age")
}
