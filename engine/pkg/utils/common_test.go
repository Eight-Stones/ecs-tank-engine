package utils

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"testing"
)

func TestCheckBitMask(t *testing.T) {
	type args struct {
		code       int
		conditions []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check simple `NotFound`",
			args: args{
				code:       common.NotFound,
				conditions: []int{common.NotFound},
			},
			want: true,
		},
		{
			name: "check complex `Fail` and `Border` and result `Fail` and `Border",
			args: args{
				code:       common.Fail | common.OkBorder,
				conditions: []int{common.Fail, common.OkBorder},
			},
			want: true,
		},
		{
			name: "check complex `Fail` and `Border` and result `Fail` and `Border\"",
			args: args{
				code:       common.Fail | common.OkBorder,
				conditions: []int{common.Ok, common.OkBorder},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckBitMask(tt.args.code, tt.args.conditions...); got != tt.want {
				t.Errorf("CheckBitMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
