package smi

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("%#v want %g got %g", test.shape, test.want, got)
		}
	}
}
