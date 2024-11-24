package main

import (
	gce "github.com/kallaurru/gocext"
	"log"
	"path/filepath"
)

func main() {
	/*	if err := Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer Close()

		x, y := TerminalDimensions()
		log.Println("parent width - ", x, "parent height - ", y)
	*/
	// banner filename
	path, err := filepath.Abs("../_assets")
	if err != nil {
		log.Fatalln(err)
	}
	gce.MakeCommonInfrastructure(path)
}
