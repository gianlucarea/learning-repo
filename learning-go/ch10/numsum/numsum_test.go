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

func TestAddNumberTypes(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		var got int = Add(10, 20)
		if got != 30 {
			t.Errorf("Add(10, 20) = %d; want 30", got)
		}
	})

	t.Run("float64", func(t *testing.T) {
		var got float64 = Add(10.5, 20.5)
		if got != 31 {
			t.Errorf("Add(10.5, 20.5) = %g; want 31", got)
		}
	})

	t.Run("uint", func(t *testing.T) {
		var got uint = Add(uint(5), uint(10))
		if got != 15 {
			t.Errorf("Add(uint(5), uint(10)) = %d; want 15", got)
		}
	})
}
