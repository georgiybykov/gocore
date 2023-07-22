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
			x1:   1.0,
			y1:   1.0,
			x2:   4.0,
			y2:   5.0,
			want: 5.0,
		},
		{
			name: "Test case #2",
			x1:   0.0,
			y1:   0.0,
			x2:   3.0,
			y2:   4.0,
			want: 5.0,
		},
		{
			name: "Test case #3",
			x1:   1.0,
			y1:   8.0,
			x2:   1.0,
			y2:   15.0,
			want: 7.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.x1, tt.y1, tt.x2, tt.y2); got != tt.want {
				t.Errorf("got distance is %v, want %v", got, tt.want)
			}
		})
	}
}
