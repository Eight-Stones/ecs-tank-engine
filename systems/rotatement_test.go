package systems

import (
	"testing"
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/v2/components"
)

func TestCanRotate(t *testing.T) {
	object := struct {
		components.Info
	}{}
	objects := []struct {
		components.Info
		components.Rotatement
		components.Position
	}{
		{
			Rotatement: components.Rotatement{
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
			Position: components.Position{},
		},
		{
			Rotatement: components.Rotatement{
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
			Position: components.Position{},
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
			name: "rotate can done",
			args: args{
				in:  &objects[0],
				now: time.Date(2000, 1, 1, 1, 1, 1, 2, time.UTC),
			},
			want: true,
		},
		{
			name: "rotate cant done",
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
			if got := CanRotate(tt.args.in, tt.args.now); got != tt.want {
				t.Errorf("CanRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetRotateDone(t *testing.T) {
	object := struct {
		components.Info
		components.Rotatement
		components.Position
	}{
		Rotatement: components.Rotatement{
			Recharge: &components.Recharge{
				DefaultDuration: time.Second * 1,
			},
		},
		Position: components.Position{},
	}
	type args struct {
		in  RotatementSystem
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
			SetRotateDone(tt.args.in, tt.args.now)
			if tt.args.in.GetRotatement().GetRecharge().Until.Unix() != tt.want.Unix() {
				t.Errorf("got = %v, want %v ", tt.args.in.GetRotatement().GetRecharge().Until, tt.want)
			}
		})
	}
}

func TestDoRotate(t *testing.T) {
	type args struct {
		direction    components.Direction
		setDirection components.Direction
	}
	tests := []struct {
		name string
		args args
		want components.Direction
	}{
		{
			name: "step right",
			args: args{
				direction:    components.Right,
				setDirection: components.Left,
			},
			want: components.Left,
		},
		{
			name: "step left",
			args: args{
				direction:    components.Left,
				setDirection: components.Up,
			},
			want: components.Up,
		},

		{
			name: "step up",
			args: args{
				direction:    components.Up,
				setDirection: components.Down,
			},
			want: components.Down,
		},
		{
			name: "step down",
			args: args{
				direction:    components.Down,
				setDirection: components.Right,
			},
			want: components.Right,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			object := struct {
				components.Rotatement
				components.Position
			}{
				Rotatement: components.Rotatement{},
				Position: components.Position{
					Direction: tt.args.direction,
				},
			}
			DoRotate(&object, tt.args.setDirection)
			if object.Direction != tt.want {
				t.Errorf("got = %s, want = %s ", object.Direction, tt.want)
			}
		})
	}
}
