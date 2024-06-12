package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectMask struct {
	container *Object
	Rect      rl.Rectangle
}

func NewRectMask(container *Object, rect rl.Rectangle) *RectMask {
	return &RectMask{
		container: container,
		Rect:      rect,
	}
}

func (rm *RectMask) OnDraw(game *Game) error {
	rl.DrawRectangleRec(rm.Rect, rl.NewColor(87, 217, 222, 100))
	return nil
}

func (rm *RectMask) OnUpdate(game *Game) error {
	rm.Rect.X = rm.container.Position.X - rm.container.Anchor.X*rm.container.Scale
	rm.Rect.Y = rm.container.Position.Y - rm.container.Anchor.Y*rm.container.Scale

	return nil
}
func (rm *RectMask) CheckCollision(other CollidableMask) bool {
	switch otherRect := other.(type) {
	case *RectMask:
		return rl.CheckCollisionRecs(rm.Rect, otherRect.Rect)
	case *CircleMask:
		return rl.CheckCollisionCircleRec(otherRect.center, otherRect.radius, rm.Rect)
	}

	return false
}
