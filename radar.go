package engine

import (
	"time"

	"ecs-tank-engine/common"
	"ecs-tank-engine/systems"
)

func (f *Field) collectRadarData(obj systems.RadarSystem) View {
	position := obj.(systems.PositionSystem)
	view := prepareView(f.gameInfo.SizeX, f.gameInfo.SizeY)

	for _, pos := range f.getAllCanPosition() {
		data := pos.GetPosition()
		view[data.X][data.Y].X = data.X
		view[data.X][data.Y].Y = data.Y
		view[data.X][data.Y].ObjectType = pos.(systems.InfoSystem).GetInfo().Type
	}

	return trimVision(position, obj.GetRadar().Radius, view, f.gameInfo.SizeX, f.gameInfo.SizeY)
}

// radar return big area around entities.Tank.
func (f *Field) radar(obj systems.InfoSystem) (int, View) {
	doing := 0b0
	now := time.Now()
	if !systems.CanRadar(obj, now) {
		return doing | common.FailRadar | common.Ban, nil
	}

	radar := obj.(systems.RadarSystem)

	systems.SetRadarDone(radar, now)

	view := f.collectRadarData(radar)

	f.cache.saveRadar(obj.GetInfo().Id, obj.GetInfo().Type, radar.GetRadar().Radius)

	return doing | common.OkRadar, view
}
