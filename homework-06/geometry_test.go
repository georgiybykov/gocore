package geometry

import (
	"testing"
)

func TestGeometry_Distance(t *testing.T) {
	tests := []struct {
		x1, y1, x2, y2 float64
		name           string
		want           float64
	}{
		{
			name: "Test case #1",
			x1:   1,
			y1:   1,
			x2:   4,
			y2:   5,
			want: 5,
		},
		{
			name: "Test case #2",
			x1:   0,
			y1:   0,
			x2:   3,
			y2:   4,
			want: 5,
		},
		{
			name: "Test case #3",
			x1:   1,
			y1:   8,
			x2:   1,
			y2:   15,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := tt.x2 - tt.x1
			y := tt.y2 - tt.y1

			if got := Distance(x, y); got != tt.want {
				t.Errorf("got distance is %v, want %v", got, tt.want)
			}
		})
	}
}
