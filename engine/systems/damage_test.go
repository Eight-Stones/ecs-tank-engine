package systems

import (
	"eight-stones/ecs-tank-engine/engine/components"
	"testing"
)

func TestCauseHitDamageSystem(t *testing.T) {
	objects := []struct {
		components.Health
		components.Damage
	}{
		{
			Health: components.Health{HitPoints: 10},
		},
		{
			Damage: components.Damage{DamagePoints: 10},
		},
		{
			Health: components.Health{HitPoints: 5},
		},
		{
			Damage: components.Damage{DamagePoints: 10},
		},
		{
			Health: components.Health{HitPoints: 5, MaxHitPoints: 12},
		},
		{
			Damage: components.Damage{DamagePoints: -10},
		},
		{
			Health: components.Health{HitPoints: 5, MaxHitPoints: 20},
		},
		{
			Damage: components.Damage{DamagePoints: -10},
		},
	}
	type args struct {
		damageTaker  HealthSystem
		damageDealer DamageSystem
	}
	tests := []struct {
		name   string
		args   args
		wantHP int
	}{
		{
			name: "10 damage 10 hp zero hit points",
			args: args{
				damageTaker:  &objects[0],
				damageDealer: &objects[1],
			},
			wantHP: 0,
		},
		{
			name: "10 damage 5 hp zero hit points",
			args: args{
				damageTaker:  &objects[2],
				damageDealer: &objects[3],
			},
			wantHP: 0,
		},
		{
			name: "-10 damage 5 hp max hit points",
			args: args{
				damageTaker:  &objects[4],
				damageDealer: &objects[5],
			},
			wantHP: 12,
		},
		{
			name: "-10 damage 5 hp 15 hit points",
			args: args{
				damageTaker:  &objects[6],
				damageDealer: &objects[7],
			},
			wantHP: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CauseHitDamageSystem(tt.args.damageTaker, tt.args.damageDealer)
			if tt.wantHP != tt.args.damageTaker.GetHealth().HitPoints {
				t.Errorf("got = %v, want %v ", tt.args.damageTaker.GetHealth().HitPoints, tt.wantHP)
			}
		})
	}
}
