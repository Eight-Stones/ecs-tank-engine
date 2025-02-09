package systems

import (
	"ecs-tank-engine/components"
	"testing"
)

func TestGetIncrementPosition(t *testing.T) {
	type args struct {
		direction components.Direction
	}
	tests := []struct {
		name  string
		args  args
		wantX int
		wantY int
	}{
		{
			name: "right",
			args: args{
				direction: components.Right,
			},
			wantX: 1,
			wantY: 0,
		},
		{
			name: "left",
			args: args{
				direction: components.Left,
			},
			wantX: -1,
			wantY: 0,
		},
		{
			name: "up",
			args: args{
				direction: components.Up,
			},
			wantX: 0,
			wantY: 1,
		},
		{
			name: "down",
			args: args{
				direction: components.Down,
			},
			wantX: 0,
			wantY: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := GetIncrementPosition(tt.args.direction)
			if gotX != tt.wantX {
				t.Errorf("GetIncrementPosition() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("GetIncrementPosition() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestChangePosition(t *testing.T) {
	type args struct {
		incX int
		incY int
		x, y int
	}
	tests := []struct {
		name         string
		args         args
		wantX, wantY int
	}{
		{
			name: "good change positive",
			args: args{
				incX: 1,
				incY: 1,
				x:    2,
				y:    1,
			},
			wantX: 3,
			wantY: 2,
		},
		{
			name: "good change negative",
			args: args{
				incX: -1,
				incY: -1,
				x:    2,
				y:    1,
			},
			wantX: 1,
			wantY: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			object := struct {
				components.Position
			}{
				Position: components.Position{
					X: tt.args.x,
					Y: tt.args.y,
				},
			}
			ChangePosition(tt.args.incX, tt.args.incY, &object)
			if object.X != tt.wantX || object.Y != tt.wantY {
				t.Errorf("got = [%d,%d], want [%d,%d] ", object.X, object.Y, tt.wantX, tt.wantY)
			}
		})
	}
}

func TestIsCollision(t *testing.T) {
	type position struct {
		components.Position
	}
	same := &position{}
	type args struct {
		first  PositionSystem
		second PositionSystem
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "collision with same",
			args: args{
				first:  same,
				second: same,
			},
			want: Fail,
		},
		{
			name: "collision with other",
			args: args{
				first: &position{
					Position: components.Position{
						X:         0,
						Y:         0,
						Direction: components.Right,
					},
				},
				second: &position{
					Position: components.Position{
						X: 1,
						Y: 0,
					},
				},
			},
			want: Success,
		},
		{
			name: "no collision with other",
			args: args{
				first: &position{
					Position: components.Position{
						X:         0,
						Y:         0,
						Direction: components.Right,
					},
				},
				second: &position{
					Position: components.Position{
						X: 2,
						Y: 0,
					},
				},
			},
			want: Fail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCollision(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("IsCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
