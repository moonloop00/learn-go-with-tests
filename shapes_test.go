package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, expectedArea: 72.0},
		{name: "Circle", shape: Circle{10}, expectedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expectedArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.expectedArea {
			t.Errorf("got %g expectedArea %g", got, tt.expectedArea)
		}
	}

}
