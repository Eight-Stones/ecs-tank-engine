package components

import "testing"

func TestShooting_HasAmmo(t *testing.T) {
	type fields struct {
		Ammo     int
		MaxAmmo  int
		Recharge Recharge
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "has ammo",
			fields: fields{
				Ammo:    1,
				MaxAmmo: 1,
			},
			want: true,
		},
		{
			name: "has no ammo",
			fields: fields{
				Ammo:    0,
				MaxAmmo: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Shooting{
				Ammo:     tt.fields.Ammo,
				MaxAmmo:  tt.fields.MaxAmmo,
				Recharge: tt.fields.Recharge,
			}
			if got := s.HasAmmo(); got != tt.want {
				t.Errorf("HasAmmo() = %v, want %v", got, tt.want)
			}
		})
	}
}
