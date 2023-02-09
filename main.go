package main

import (
	"botkiosgamercodm/handler/rest"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	rest.StartApp()
}
