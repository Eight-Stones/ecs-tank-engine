package systems

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"reflect"
	"testing"
)

func TestAddAction(t *testing.T) {
	object := struct {
		components.Info
		components.Statistic
	}{
		Info:      components.Info{},
		Statistic: components.Statistic{},
	}
	type args struct {
		code int
		in   InfoSystem
	}
	tests := []struct {
		name string
		args args

		want []int
	}{
		{
			name: "good added stat",
			args: args{
				code: 200,
				in:   &object,
			},
			want: []int{200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddAction(tt.args.code, tt.args.in)
			if !reflect.DeepEqual(tt.args.in.(StatisticSystem).GetStatistic().Actions, tt.want) {
				t.Errorf("got= %v, want %v", tt.args.in.(StatisticSystem).GetStatistic().Actions, tt.want)
			}
		})
	}
}
