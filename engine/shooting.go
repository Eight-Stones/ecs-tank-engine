package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/entities"
	"eight-stones/ecs-tank-engine/engine/pkg/utils"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

func (f *Field) shoot(id string) int {
	obj, code := f.find(id)
	doing := 0b0 | code
	if utils.CheckBitMask(code, common.NotFound) {
		return doing | common.FailShot
	}

	now := time.Now()
	if !systems.CanShoot(obj, now) {
		return doing | common.FailShot | common.Ban
	}

	shooting := obj.(systems.ShootingSystem)

	systems.SetShotDone(shooting, now)

	doing = doing | f.move(f.createBullet(obj), now)
	if utils.CheckBitMask(doing, common.FailStep) {
		return doing | common.FailShot
	}

	return doing | common.OkShot
}

func (f *Field) createBullet(in systems.CommonSystem) systems.CommonSystem {
	movement := in.(systems.RotatementSystem)
	// TODO переосмыслить этот момент, кажется, что система позиционирования должна иметь направление
	bullet := entities.NewBullet(
		&f.cfg.Bullet,
		in.GetCommon(),
		movement.GetPosition().X,
		movement.GetPosition().Y,
		movement.GetMovement().Direction,
	)
	f.Objects = append(f.Objects, &bullet)
	return &bullet
}
