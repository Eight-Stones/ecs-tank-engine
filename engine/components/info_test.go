package components

import "testing"

func TestObjectType_String(t *testing.T) {
	tests := []struct {
		name string
		ot   ObjectType
		want string
	}{
		{
			name: "empty",
			ot:   0,
			want: "",
		},
		{
			name: "tank",
			ot:   TypeTankId,
			want: "tank",
		},
		{
			name: "bullet",
			ot:   TypeBulletId,
			want: "bullet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ot.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
