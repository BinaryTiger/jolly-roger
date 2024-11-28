/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"binarytiger/jolly_roger/cmd"
	"database/sql"
	"fmt"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func main() {
	//saveRawHook()
	cmd.Execute()
}

func saveRawHook() {
	var version string
	db, _ := sql.Open("sqlite3", "file:local.db") // #TODO load as config
	db.QueryRow(`SELECT sqlite_version()`).Scan(&version)
	fmt.Println(version)
}
