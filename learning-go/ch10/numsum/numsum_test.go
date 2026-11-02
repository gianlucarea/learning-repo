package numsum

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "two positive numbers", a: 2, b: 3, want: 5},
		{name: "two negative numbers", a: -2, b: -3, want: -5},
		{name: "zero", a: 0, b: 0, want: 0},
		{name: "positive and negative number", a: 5, b: -3, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
