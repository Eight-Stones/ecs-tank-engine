package systems

import (
	"testing"
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/components"
)

func TestCanRadar(t *testing.T) {
	object := struct {
		components.Info
	}{}
	objects := []struct {
		components.Info
		components.Radar
		components.Position
	}{
		{
			Radar: components.Radar{
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
		},
		{
			Radar: components.Radar{
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
			name: "radar can done",
			args: args{
				in:  &objects[0],
				now: time.Date(2000, 1, 1, 1, 1, 1, 2, time.UTC),
			},
			want: true,
		},
		{
			name: "radar cant done",
			args: args{
				in:  &objects[1],
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
			if got := CanRadar(tt.args.in, tt.args.now); got != tt.want {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSetRadarDone(t *testing.T) {
	object := struct {
		components.Info
		components.Radar
		components.Position
	}{
		Radar: components.Radar{
			Recharge: &components.Recharge{
				DefaultDuration: time.Second * 1,
			},
		},
		Position: components.Position{},
	}
	type args struct {
		in  RadarSystem
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
			SetRadarDone(tt.args.in, tt.args.now)
			if tt.args.in.GetRadar().GetRecharge().Until.Unix() != tt.want.Unix() {
				t.Errorf("got = %v, want %v ", tt.args.in.GetRadar().GetRecharge().Until, tt.want)
			}
		})
	}
}
