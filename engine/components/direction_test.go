package components

import "testing"

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name string
		d    Direction
		want string
	}{
		{
			name: "left",
			d:    Left,
			want: "left",
		},
		{
			name: "right",
			d:    Right,
			want: "right",
		},
		{
			name: "up",
			d:    Up,
			want: "up",
		},
		{
			name: "down",
			d:    Down,
			want: "down",
		},
		{
			name: "empty",
			d:    0,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
