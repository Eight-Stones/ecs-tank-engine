package components

import (
	"reflect"
	"testing"
)

func TestStepMoveSystem(t *testing.T) {
	type args struct {
		direction uint
		in        *Tank
		want      *Tank
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "GOOD",
			args: args{
				direction: Left,
				in: &Tank{
					Info:     Info{Id: "test"},
					Position: Position{X: 1, Y: 1},
					Movement: Movement{Direction: Right, AutoMove: false},
				},
				want: &Tank{
					Info:     Info{Id: "test"},
					Position: Position{X: 1, Y: 1},
					Movement: Movement{Direction: Left, AutoMove: false},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StepMoveSystem(tt.args.direction, tt.args.in)
			reflect.DeepEqual(tt.args.in, tt.args.want)
		})
	}
}
