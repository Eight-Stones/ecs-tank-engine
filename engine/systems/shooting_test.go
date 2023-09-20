package systems

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"testing"
	"time"
)

func TestCanShoot(t *testing.T) {
	object := struct {
		components.Info
	}{}
	objects := []struct {
		components.Info
		components.Shooting
		components.Position
	}{
		{
			Shooting: components.Shooting{
				Ammo: 1,
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
		},
		{
			Shooting: components.Shooting{
				Ammo: 0,
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
		},
		{
			Shooting: components.Shooting{
				Ammo: 1,
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
		},
		{
			Shooting: components.Shooting{
				Ammo: 0,
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
		},
	}
	type args struct {
		in  InfoSystem
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "shoot recharge done with ammo",
			args: args{
				in:  &objects[0],
				now: time.Date(2000, 1, 1, 1, 1, 1, 2, time.UTC),
			},
			want: true,
		},
		{
			name: "shoot recharge done without ammo",
			args: args{
				in:  &objects[1],
				now: time.Date(2000, 1, 1, 1, 1, 1, 2, time.UTC),
			},
			want: false,
		},
		{
			name: "shoot recharge does not done with ammo",
			args: args{
				in:  &objects[2],
				now: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "shoot recharge does not done without ammo",
			args: args{
				in:  &objects[3],
				now: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "without implementations",
			args: args{
				in:  &object,
				now: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanShoot(tt.args.in, tt.args.now); got != tt.want {
				t.Errorf("CanShoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetShotDone(t *testing.T) {
	object := struct {
		components.Info
		components.Shooting
		components.Position
	}{
		Shooting: components.Shooting{
			Recharge: &components.Recharge{
				DefaultDuration: time.Second * 1,
			},
		},
		Position: components.Position{},
	}
	type args struct {
		in  ShootingSystem
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "good set recharge",
			args: args{
				in:  &object,
				now: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
			},
			want: time.Date(2000, 1, 1, 1, 1, 2, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetShotDone(tt.args.in, tt.args.now)
			if tt.args.in.GetShooting().GetRecharge().Until.Unix() != tt.want.Unix() {
				t.Errorf("got = %v, want %v ", tt.args.in.GetShooting().GetRecharge().Until, tt.want)
			}
		})
	}
}
