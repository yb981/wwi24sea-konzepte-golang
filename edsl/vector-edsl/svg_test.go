package main

import (
	"os"
	"testing"
)

func TestSVGToSVG(t *testing.T) {
	s := svg{
		width:  100,
		height: 100,
		content: []Element{
			rect{width: 50, height: 50, x: 10, y: 10, rx: 5, ry: 5, fill: "red"},
			circle{r: 20, cx: 30, cy: 30, fill: "blue"},
			line{x1: 10, y1: 10, x2: 50, y2: 50, style: "stroke:black;"},
			text{x: 20, y: 40, fill: "black", content: "Hello"},
			ellipse{rx: 25, ry: 15, cx: 50, cy: 50, style: "fill:green;"},
		},
	}

	expected := `<svg width="100" height="100" xmlns="http://www.w3.org/2000/svg">
	<rect width="50" height="50" x="10" y="10" rx="5" ry="5" fill="red" />
	<circle r="20" cx="30" cy="30" fill="blue" />
	<line x1="10" y1="10" x2="50" y2="50" style="stroke:black;" />
	<text x="20" y="40" fill="black">Hello</text>
	<ellipse rx="25" ry="15" cx="50" cy="50" style="fill:green;"
</svg>`

	if s.toSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, s.toSVG())
	}
}

func TestSVGSaveSVG(t *testing.T) {
	s := svg{width: 100, height: 100, content: []Element{}}
	s.saveSVG("test")
	if _, err := os.Stat("test.svg"); os.IsNotExist(err) {
		t.Errorf("File test.svg was not created")
	}
	_ = os.Remove("test.svg")
}

func TestRectToSVG(t *testing.T) {
	r := rect{width: 50, height: 50, x: 10, y: 10, rx: 5, ry: 5, fill: "red"}
	expected := `	<rect width="50" height="50" x="10" y="10" rx="5" ry="5" fill="red" />`
	if r.toSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, r.toSVG())
	}
}

func TestCircleToSVG(t *testing.T) {
	c := circle{r: 20, cx: 30, cy: 30, fill: "blue"}
	expected := `	<circle r="20" cx="30" cy="30" fill="blue" />`
	if c.toSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, c.toSVG())
	}
}

func TestLineToSVG(t *testing.T) {
	l := line{x1: 10, y1: 10, x2: 50, y2: 50, style: "stroke:black;"}
	expected := `	<line x1="10" y1="10" x2="50" y2="50" style="stroke:black;" />`
	if l.toSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, l.toSVG())
	}
}

func TestTextToSVG(t *testing.T) {
	txt := text{x: 20, y: 40, fill: "black", content: "Hello"}
	expected := `	<text x="20" y="40" fill="black">Hello</text>`
	if txt.toSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, txt.toSVG())
	}
}

func TestEllipseToSVG(t *testing.T) {
	e := ellipse{rx: 25, ry: 15, cx: 50, cy: 50, style: "fill:green;"}
	expected := `	<ellipse rx="25" ry="15" cx="50" cy="50" style="fill:green;"`
	if e.toSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, e.toSVG())
	}
}