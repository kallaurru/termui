package main

import (
	"image"
	"log"

	. "github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/widgets"
)

func main() {
	if err := Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer Close()

	w, h := TerminalDimensions()
	log.Println("width - ", w, "height - ", h)

	parent := image.Rect(0, 0, w, h)
	title := "Information Box"
	center := image.Pt(parent.Dx()/2, parent.Dy()/2)
	db := widgets.NewDialogBox(center, title, "Test information message. Period is correct")

	Render(db)
	uiEvents := PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Enter>":
				lab := "Taked"
				db.SetButtonLabel(lab)
				Render(db)

			}
		}
	}
}
