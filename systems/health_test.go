package systems

import (
	"ecs-tank-engine/components"
	"testing"
)

func TestIsAliveHealthSystem(t *testing.T) {
	objects := []struct {
		components.Health
	}{
		{
			Health: components.Health{HitPoints: 10},
		},
		{
			Health: components.Health{HitPoints: 0},
		},
	}
	type args struct {
		in HealthSystem
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is alive",
			args: args{
				in: &objects[0],
			},
			want: true,
		},
		{
			name: "is dead",
			args: args{
				in: &objects[1],
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAliveHealthSystem(tt.args.in); got != tt.want {
				t.Errorf("IsAliveHealthSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChangeHPLevelHealthSystem(t *testing.T) {
	objects := []struct {
		components.Health
	}{
		{
			Health: components.Health{HitPoints: 10},
		},
		{
			Health: components.Health{HitPoints: 5},
		},
		{
			Health: components.Health{HitPoints: 5, MaxHitPoints: 12},
		},
		{
			Health: components.Health{HitPoints: 5, MaxHitPoints: 20},
		},
	}
	type args struct {
		in    HealthSystem
		delta int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "set hp to 0",
			args: args{
				in:    &objects[0],
				delta: 10,
			},
			want: 0,
		},
		{
			name: "set hp to 0",
			args: args{
				in:    &objects[1],
				delta: 10,
			},
			want: 0,
		},
		{
			name: "set hp to maximum hp",
			args: args{
				in:    &objects[2],
				delta: -10,
			},
			want: 12,
		},
		{
			name: "set hp between maximum and start value",
			args: args{
				in:    &objects[3],
				delta: -10,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChangeHPLevelHealthSystem(tt.args.in, tt.args.delta)
			if tt.want != tt.args.in.GetHealth().HitPoints {
				t.Errorf("got = %v, want %v ", tt.args.in.GetHealth().HitPoints, tt.want)
			}
		})
	}
}

func TestDisappear(t *testing.T) {
	object := struct {
		components.Health
	}{
		Health: components.Health{HitPoints: 999999},
	}

	type args struct {
		in HealthSystem
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "remove all hit points",
			args: args{
				in: &object,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Disappear(tt.args.in)
			if tt.want != tt.args.in.GetHealth().HitPoints {
				t.Errorf("got = %v, want %v ", tt.args.in.GetHealth().HitPoints, tt.want)
			}
		})
	}
}
