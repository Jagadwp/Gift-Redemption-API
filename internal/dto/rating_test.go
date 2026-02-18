package dto

import "testing"

func TestRoundToHalf(t *testing.T) {
	tests := []struct {
		name string
		in   float64
		want float64
	}{
		{"round 3.2 to 3.0", 3.2, 3.0},
		{"round 3.6 to 3.5", 3.6, 3.5},
		{"round 3.9 to 4.0", 3.9, 4.0},
		{"round 4.25 to 4.0", 4.25, 4.0},
		{"round 4.26 to 4.5", 4.26, 4.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundToHalf(tt.in); got != tt.want {
				t.Fatalf("RoundToHalf(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
