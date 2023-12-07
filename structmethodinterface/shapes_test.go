package structmethodinterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := Perimeter(rectangle)
	expected := 40.0

	if actual != expected {
		t.Errorf("expected %.2f but actual %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{5, 3}, 7.5},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()

		if actual != tt.expected {
			t.Errorf("%#v expected %.2f but actual %.2f", tt.shape, tt.expected, actual)
		}
	}
}
