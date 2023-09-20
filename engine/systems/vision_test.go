package systems

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"testing"
	"time"
)

func TestCanVision(t *testing.T) {
	object := struct {
		components.Info
	}{}
	objects := []struct {
		components.Info
		components.Vision
		components.Position
	}{
		{
			Vision: components.Vision{
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
		},
		{
			Vision: components.Vision{
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
			name: "vision can done",
			args: args{
				in:  &objects[0],
				now: time.Date(2000, 1, 1, 1, 1, 1, 2, time.UTC),
			},
			want: true,
		},
		{
			name: "vision cant done",
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
			if got := CanVision(tt.args.in, tt.args.now); got != tt.want {
				t.Errorf("CanVision() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetVisionDone(t *testing.T) {
	object := struct {
		components.Info
		components.Vision
		components.Position
	}{
		Vision: components.Vision{
			Recharge: &components.Recharge{
				DefaultDuration: time.Second * 1,
			},
		},
		Position: components.Position{},
	}
	type args struct {
		in  VisionSystem
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "good set vision",
			args: args{
				in:  &object,
				now: time.Date(2000, 1, 1, 1, 1, 1, 0, time.UTC),
			},
			want: time.Date(2000, 1, 1, 1, 1, 2, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetVisionDone(tt.args.in, tt.args.now)
			if tt.args.in.GetVision().GetRecharge().Until.Unix() != tt.want.Unix() {
				t.Errorf("got = %v, want %v ", tt.args.in.GetVision().GetRecharge().Until, tt.want)
			}
		})
	}
}
