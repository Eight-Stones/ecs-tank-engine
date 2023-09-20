package engine

import (
	"eight-stones/ecs-tank-engine/engine/common"
	"eight-stones/ecs-tank-engine/engine/systems"
	"time"
)

// Cell ...
type Cell struct {
	X          int
	Y          int
	ObjectType string
}

// View ...
type View [][]Cell

// prepareView prepared empty view.
func prepareView(x, y int) View {
	view := make(View, y-1)
	for idx := range view {
		view[idx] = make([]Cell, x-1)
	}
	return view
}

// trimVision trims vision based on radius of vision.
func trimVision(object systems.PositionSystem, radius int, view View, mX, mY int) View {
	x, y := object.GetPosition().X, object.GetPosition().Y

	mX = mX - 1
	mY = mY - 1

	left := x - radius
	right := x + radius
	top := y + radius
	bottom := y - radius

	if left < 0 {
		left = 0
	}
	if bottom < 0 {
		bottom = 0
	}
	if right > mX {
		right = mX
	}
	if top > mY {
		top = mY
	}

	newView := make(View, 0, radius*2)
	for idy := bottom; idy <= top; idy++ {
		newView = append(newView, view[idy][left:right+1])
	}

	return newView
}

// collectVisionData ...
func (f *Field) collectVisionData(obj systems.VisionSystem) View {
	position := obj.(systems.PositionSystem)
	view := prepareView(f.gameInfo.SizeX, f.gameInfo.SizeY)

	for _, pos := range f.getAllCanPosition() {
		data := pos.GetPosition()
		view[data.X][data.Y].X = data.X
		view[data.X][data.Y].Y = data.Y
		view[data.X][data.Y].ObjectType = defineType(pos.(systems.InfoSystem))
	}

	return trimVision(position, obj.GetVision().Radius, view, f.gameInfo.SizeX, f.gameInfo.SizeY)
}

// vision return small area around entities.Tank.
func (f *Field) vision(obj systems.InfoSystem) (int, View) {
	doing := 0b0
	now := time.Now()
	if !systems.CanVision(obj, now) {
		return doing | common.FailVision | common.Ban, nil
	}

	vision := obj.(systems.VisionSystem)

	systems.SetVisionDone(vision, now)

	view := f.collectVisionData(vision)

	f.cache.saveVision(obj.GetInfo().Id, obj.GetInfo().Type, vision.GetVision().Radius)

	return doing | common.OkVision, view
}
