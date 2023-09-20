package components

import (
	"testing"
	"time"
)

func TestRecharge_IsRechargeDone(t *testing.T) {
	type fields struct {
		Until           time.Time
		DefaultDuration time.Duration
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "recharge end",
			fields: fields{
				Until:           time.Now().Add(time.Second * -1),
				DefaultDuration: 0,
			},
			args: args{
				now: time.Now(),
			},
			want: true,
		},
		{
			name: "recharge continue",
			fields: fields{
				Until:           time.Now(),
				DefaultDuration: 0,
			},
			args: args{
				now: time.Now().Add(time.Second * -1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Recharge{
				Until:           tt.fields.Until,
				DefaultDuration: tt.fields.DefaultDuration,
			}
			if got := r.IsRechargeDone(tt.args.now); got != tt.want {
				t.Errorf("IsRechargeDone() = %v, want %v", got, tt.want)
			}
		})
	}
}
