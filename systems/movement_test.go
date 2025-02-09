package systems

import (
	"testing"
	"time"

	"github.com/Eight-Stones/ecs-tank-engine/components"
)

func TestCanStep(t *testing.T) {
	object := struct {
		components.Info
	}{}
	objects := []struct {
		components.Info
		components.Movement
		components.Position
	}{
		{
			Movement: components.Movement{
				Recharge: &components.Recharge{
					Until: time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC),
				},
			},
			Position: components.Position{},
		},
		{
			Movement: components.Movement{
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
			name: "step can done",
			args: args{
				in:  &objects[0],
				now: time.Date(2000, 1, 1, 1, 1, 1, 2, time.UTC),
			},
			want: true,
		},
		{
			name: "step cant done",
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
			if got := CanStep(tt.args.in, tt.args.now); got != tt.want {
				t.Errorf("CanStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetStepDone(t *testing.T) {
	object := struct {
		components.Info
		components.Movement
		components.Position
	}{
		Movement: components.Movement{
			Recharge: &components.Recharge{
				DefaultDuration: time.Second * 1,
			},
		},
		Position: components.Position{},
	}
	type args struct {
		in  MovementSystem
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
			SetStepDone(tt.args.in, tt.args.now)
			if tt.args.in.GetMovement().GetRecharge().Until.Unix() != tt.want.Unix() {
				t.Errorf("got = %v, want %v ", tt.args.in.GetMovement().GetRecharge().Until, tt.want)
			}
		})
	}
}

func TestDoStep(t *testing.T) {
	type args struct {
		x, y      int
		direction components.Direction
	}
	tests := []struct {
		name         string
		args         args
		wantX, wantY int
	}{
		{
			name: "step right",
			args: args{
				x:         1,
				y:         1,
				direction: components.Right,
			},
			wantX: 2,
			wantY: 1,
		},
		{
			name: "step left",
			args: args{
				x:         1,
				y:         1,
				direction: components.Left,
			},
			wantX: 0,
			wantY: 1,
		},

		{
			name: "step up",
			args: args{
				x:         1,
				y:         1,
				direction: components.Up,
			},
			wantX: 1,
			wantY: 2,
		},
		{
			name: "step down",
			args: args{
				x:         1,
				y:         1,
				direction: components.Down,
			},
			wantX: 1,
			wantY: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			object := struct {
				components.Movement
				components.Position
			}{
				Movement: components.Movement{},
				Position: components.Position{
					X:         tt.args.x,
					Y:         tt.args.y,
					Direction: tt.args.direction,
				},
			}
			DoStep(&object)
			if object.X != tt.wantX || object.Y != tt.wantY {
				t.Errorf("got = [%d,%d], want [%d,%d] ", object.X, object.Y, tt.wantX, tt.wantY)
			}
		})
	}
}
