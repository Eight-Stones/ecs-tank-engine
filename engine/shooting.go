package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

// shoot shoots from tank.
func (f *Field) shoot(obj systems.InfoSystem) int {
	doing := 0b0
	now := time.Now()
	if !systems.CanShoot(obj, now) {
		return doing | common.FailShot | common.Ban
	}

	shooting := obj.(systems.ShootingSystem)

	systems.SetShotDone(shooting, now)

	f.cache.saveShoot(
		obj.GetInfo().Id,
		shooting.GetShooting().Ammo,
		obj.(systems.PositionSystem).GetPosition().Direction,
	)

	bullet := f.createBullet(obj)

	doing = doing | f.move(bullet, now)
	if utils.CheckBitMask(doing, common.FailStep) {
		return doing | common.OkShot
	}

	f.cache.saveCreate(
		bullet.GetInfo().Id,
		bullet.(systems.PositionSystem).GetPosition().Direction,
		[]int{bullet.(systems.PositionSystem).GetPosition().X, bullet.(systems.PositionSystem).GetPosition().Y},
		bullet.(systems.HealthSystem).GetHealth().HitPoints,
	)

	return doing | common.OkShot
}

// createBullet created bullet object.
func (f *Field) createBullet(in systems.InfoSystem) systems.InfoSystem {
	position := in.(systems.PositionSystem)
	bullet := entities.NewBullet(
		&f.cfg.Bullet,
		in.GetInfo(),
		position.GetPosition().X,
		position.GetPosition().Y,
		position.GetPosition().Direction,
	)
	f.Objects = append(f.Objects, &bullet)
	return &bullet
}
