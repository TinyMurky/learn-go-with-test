package shape

import "testing"

// func TestPerimeter(t *testing.T) {
// 	got := Perimeter(10.0, 10.0)
// 	want := 40.0
//
// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				Width:  12,
				Height: 6,
			},
			hasArea: 72.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Radius: 10,
			},
			hasArea: 314.1592653589793,
		},
		{
			name: "Triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			hasArea: 36.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name      string
		shape     Shape
		hasPerim  float64
		wantPanic bool
	}{
		{
			name:     "Rectangle",
			shape:    Rectangle{Width: 12, Height: 6},
			hasPerim: 36.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Radius: 10,
			},
			hasPerim: 62.83185307179586,
		},
		{
			name: "Triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
				Sides:  []float64{5, 5, 6},
			},
			hasPerim: 16.0,
		},
		{
			name: "Triangle with wrong sides",
			shape: Triangle{
				Base:   12,
				Height: 6,
				Sides:  []float64{5, 6, 7, 8}, // More than 3 sides, should cause a panic
			},
			wantPanic: true,
		},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if we expect a panic
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic for %s, but did not panic", tt.name)
					}
				}()
			}

			got := tt.shape.Perimeter()
			if got != tt.hasPerim {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerim)
			}
		})
	}
}
