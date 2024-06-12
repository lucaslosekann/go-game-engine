package engine

type CollidableMask interface {
	CheckCollision(other CollidableMask) bool
	OnDraw(game *Game) error
	OnUpdate(game *Game) error
}
