package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

func (f *Field) collectRadarData(obj systems.RadarSystem) View {
	position := obj.(systems.PositionSystem)
	view := prepareView(f.metaInfo.SizeX, f.metaInfo.SizeY)

	for _, pos := range f.getAllCanPosition() {
		data := pos.GetPosition()
		view[data.X][data.Y].X = data.X
		view[data.X][data.Y].Y = data.Y
		view[data.X][data.Y].ObjectType = DefineType(pos.(systems.InfoSystem))
	}

	return trimVision(position, obj.GetRadar().Radius, view, f.metaInfo.SizeX, f.metaInfo.SizeY)
}

// vision return small area around entities.Tank.
func (f *Field) radar(obj systems.InfoSystem) (int, View) {
	doing := 0b0
	now := time.Now()
	if !systems.CanVision(obj, now) {
		return doing | common.FailRadar | common.Ban, nil
	}

	radar := obj.(systems.RadarSystem)

	systems.SetRadarDone(radar, now)

	view := f.collectRadarData(radar)

	return doing | common.OkRadar, view
}
