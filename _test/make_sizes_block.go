package main

import (
	"image"
	"log"
)

func main() {
	parent := image.Rect(0, 0, 70, 23)
	str := "Idslkjfsklfdj skdlfjskfjslkfj8elskfdj skdfjsfjwesljfklsjf"
	limitStr := 40
	lenStr := len(str)
	countRows := lenStr / limitStr
	if lenStr%limitStr > 0 {
		countRows++
	}
	diffW := parent.Dx() - lenStr
	shift := diffW / 2
	log.Println("weight - ", diffW, "rows - ", countRows, "shift - ", shift)
}
