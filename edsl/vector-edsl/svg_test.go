package svg

import (
	"os"
	"testing"
)

func TestSVGToSVG(t *testing.T) {
	s := Svg{
		Width:  100,
		Height: 100,
		Content: []Element{
			Rect{Width: 50, Height: 50, X: 10, Y: 10, Rx: 5, Ry: 5, Fill: "red"},
			Circle{R: 20, Cx: 30, Cy: 30, Fill: "blue"},
			Line{X1: 10, Y1: 10, X2: 50, Y2: 50, Style: "stroke:black;"},
			Text{X: 20, Y: 40, Fill: "black", Content: "Hello", Size: 12},
			Ellipse{Rx: 25, Ry: 15, Cx: 50, Cy: 50, Style: "fill:green;"},
		},
	}

	expected := `<svg width="100" height="100" xmlns="http://www.w3.org/2000/svg">
	<rect width="50" height="50" x="10" y="10" rx="5" ry="5" fill="red" />
	<circle r="20" cx="30" cy="30" fill="blue" />
	<line x1="10" y1="10" x2="50" y2="50" style="stroke:black;" />
	<text x="20" y="40" fill="black" font-size="12">Hello</text>
	<ellipse rx="25" ry="15" cx="50" cy="50" style="fill:green;" />
</svg>`

	if s.ToSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, s.ToSVG())
	}
}

func TestSVGSaveSVG(t *testing.T) {
	s := Svg{Width: 100, Height: 100, Content: []Element{}}
	s.SaveSVG("test")
	if _, err := os.Stat("test.svg"); os.IsNotExist(err) {
		t.Errorf("File test.svg was not created")
	}
	_ = os.Remove("test.svg")
}

func TestRectToSVG(t *testing.T) {
	r := Rect{Width: 50, Height: 50, X: 10, Y: 10, Rx: 5, Ry: 5, Fill: "red"}
	expected := `	<rect width="50" height="50" x="10" y="10" rx="5" ry="5" fill="red" />`
	if r.ToSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, r.ToSVG())
	}
}

func TestCircleToSVG(t *testing.T) {
	c := Circle{R: 20, Cx: 30, Cy: 30, Fill: "blue"}
	expected := `	<circle r="20" cx="30" cy="30" fill="blue" />`
	if c.ToSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, c.ToSVG())
	}
}

func TestLineToSVG(t *testing.T) {
	l := Line{X1: 10, Y1: 10, X2: 50, Y2: 50, Style: "stroke:black;"}
	expected := `	<line x1="10" y1="10" x2="50" y2="50" style="stroke:black;" />`
	if l.ToSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, l.ToSVG())
	}
}

func TestTextToSVG(t *testing.T) {
	txt := Text{X: 20, Y: 40, Fill: "black", Content: "Hello", Size: 12}
	expected := `	<text x="20" y="40" fill="black" font-size="12">Hello</text>`
	if txt.ToSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, txt.ToSVG())
	}
}

func TestEllipseToSVG(t *testing.T) {
	e := Ellipse{Rx: 25, Ry: 15, Cx: 50, Cy: 50, Style: "fill:green;"}
	expected := `	<ellipse rx="25" ry="15" cx="50" cy="50" style="fill:green;" />`
	if e.ToSVG() != expected {
		t.Errorf("Expected %v, got %v", expected, e.ToSVG())
	}
}
