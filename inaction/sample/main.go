package main

import (
	"log"
	"os"

	_ "./matchers"
	"./search"
)

func init() {
	log.SetOutput(os.Stdout) // todo why
}

func main() {
	search.Run("president")
}
