package engine

type Component interface {
	OnUpdate(game *Game) error
	OnDraw(game *Game) error
}
