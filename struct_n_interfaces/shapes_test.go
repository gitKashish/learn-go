package structninterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("find perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Height: 12, Width: 6}, want: 72.0},
		{name: "circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "want", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want
			if got != want {
				t.Errorf("got %g want %g", got, want)
			}
		})
	}
}
