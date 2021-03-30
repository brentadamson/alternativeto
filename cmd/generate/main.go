package main

import (
	"alternativeto"
	"log"
)

// This generates the README.md files
func main() {
	// Generate the main README.md. Might have to just link to other READMEs if it gets too unwieldy
	if err := alternativeto.GenerateMainREADME(); err != nil {
		log.Fatalln(err)
	}
}
