package main

import (
	. "github.com/kallaurru/termui/v3"
	"github.com/kallaurru/termui/v3/widgets"
	"image"
	"log"
)

func main() {
	if err := Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer Close()

	w, h := TerminalDimensions()
	log.Println("width - ", w, "height - ", h)

	center := image.Pt(w/2, h/2)
	fr := widgets.NewFrame(center, "Input the number")

	Render(fr)
	uiEvents := PollEvents()
loop:

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "<F10>":
				return
			case "<Enter>":
				lab := "Taked"
				fr.Block.MakeGlamourTitle(lab)
				fr.CaptureModeOff()
				Render(fr)
			case "<Insert>":
				fr.Block.MakeGlamourTitle("Input mode on")
				fr.CaptureModeOn()
				Render(fr)
				goto loop
			default:
				fr.AddSymbol(e.ID, e.Ch)
				Render(fr)
			}
		}
	}
}
