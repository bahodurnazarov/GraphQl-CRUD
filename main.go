package main

import "os"

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDatabase()
	createTable()

	defer databasa.DBConn.Close()

	srv :=
}