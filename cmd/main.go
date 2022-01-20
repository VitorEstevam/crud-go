/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	// "github.com/VitorEstevam/crud-go/cmd/cmd"
	"fmt"

	"github.com/VitorEstevam/crud-go/database"
)

func main() {
	r, err := database.NewRepository("sqlite3", "./data.db", 10, 10)
	if err != nil {
		fmt.Println("err", err)
	}
	r.CreateTable()
	ev := database.EventModel{Name: "aaaa", Year: 1123}
	err = r.Create(ev)
	if err != nil {
		fmt.Println("err", err)
	}
}
