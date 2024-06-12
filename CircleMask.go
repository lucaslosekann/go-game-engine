package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CircleMask struct {
	container *Object
	center    rl.Vector2
	radius    float32
}

func NewCircleMask(container *Object, center rl.Vector2, radius float32) *CircleMask {
	return &CircleMask{
		container: container,
		center:    center,
		radius:    radius,
	}
}

func (rm *CircleMask) OnDraw(game *Game) error {
	rl.DrawCircleV(rm.center, rm.radius, rl.NewColor(87, 217, 222, 100))
	return nil

}

func (rm *CircleMask) OnUpdate(game *Game) error {
	rm.center = rl.NewVector2(
		rm.container.Position.X-rm.container.Anchor.X*rm.container.Scale+rm.radius,
		rm.container.Position.Y-rm.container.Anchor.Y*rm.container.Scale+rm.radius,
	)
	return nil
}

func (rm *CircleMask) CheckCollision(other CollidableMask) bool {
	switch otherRect := other.(type) {
	case *RectMask:
		return rl.CheckCollisionCircleRec(rm.center, rm.radius, otherRect.Rect)
	case *CircleMask:
		return rl.CheckCollisionCircles(rm.center, rm.radius, otherRect.center, otherRect.radius)
	}

	return false
}
