package main

import (
	"math/rand"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
)

func main() {
	f, err := os.Create("hello.svg")
	if err != nil {
		panic(err)
	}
	canvas := svg.New(f)
	data := []struct {
		Month string
		Usage int
	}{
		{"Jan", 0},
		{"Feb", 0},
		{"Mar", 0},
		{"Apr", 0},
		{"May", 0},
		{"Jun", 0},
		{"Jul", 0},
		{"Aug", 0},
		{"Sep", 0},
		{"Oct", 0},
		{"Nov", 0},
		{"Dec", 0},
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < len(data); i++ {
		data[i].Usage = r.Intn(250)
	}

	width := len(data)*60 + 10
	height := 250
	threshold := 120
	max := 0
	for _, item := range data {
		if item.Usage > max {
			max = item.Usage
		}
	}
	canvas.Start(width, height+1000)
	for i, val := range data {
		percent := val.Usage * (height - 50) / max
		canvas.Rect(i*60+10, height-50-percent, 50, percent, "fill:rgb(77,200,232)")
		canvas.Text(i*60+33, height-30, val.Month, "text-anchor:middle;font-size:20px;fill:rgb(0,0,0)")
	}
	threshPercent := threshold * (height - 50) / max
	canvas.Line(0, height-50-threshPercent, width, height-50-threshPercent, "stroke:rgb(250,120,120);stroke-width:2px;opacity:0.8")
	canvas.Rect(0, 0, width, height-50-threshPercent, "fill:rgb(250,120,120);opacity:0.3")
	canvas.Text(width/2, height/2-40, "Critical consumption", "text-anchor:middle;font-size:30px;fill:black")
	canvas.Line(0, height-50, width, height-50, "stroke:rgb(0,0,0);stroke-width:3px")
	canvas.Text(width/2, height, "Energy consumption", "text-anchor:middle;font-size:20px;fill:rgb(250,0,0);font-family: cursive;")
	//stickman
	shiftY := 200
	// Draw the head
	canvas.Circle(150, 100+shiftY, 50, "fill:none;stroke:black;stroke-width:2")

	// Draw the eyes
	canvas.Circle(120, 90+shiftY, 10, "fill:black")
	canvas.Circle(180, 90+shiftY, 10, "fill:black")

	// Draw the smile
	canvas.Arc(130, 130+shiftY, 120, 120, 120+shiftY, false, false, 170, 130+shiftY, "stroke:black;stroke-width:2")

	// Draw the body
	canvas.Line(150, 150+shiftY, 150, 250+shiftY, "stroke:black;stroke-width:2")

	// Draw the arms
	canvas.Line(150, 180+shiftY, 200, 220+shiftY, "stroke:black;stroke-width:2")
	canvas.Line(150, 180+shiftY, 100, 220+shiftY, "stroke:black;stroke-width:2")

	// Draw the legs
	canvas.Line(150, 250+shiftY, 200, 300+shiftY, "stroke:black;stroke-width:2")
	canvas.Line(150, 250+shiftY, 100, 300+shiftY, "stroke:black;stroke-width:2")

	// Draw the gun
	canvas.Rect(220, 215+shiftY, 20, 50, "fill:black")
	canvas.Rect(200, 215+shiftY, 20, 15, "fill:black")

	canvas.End()

}
