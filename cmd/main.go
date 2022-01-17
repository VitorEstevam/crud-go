/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/VitorEstevam/crud-go/cmd/cmd"
	"github.com/VitorEstevam/crud-go/database"
)

func main() {
	cmd.Execute()
	db := database.Connect()
	db.Close()
}
