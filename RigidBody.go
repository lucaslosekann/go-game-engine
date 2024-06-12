package engine

type Collidable struct {
	Container   *Object
	Mask        CollidableMask
	onCollision []func(obj *Object)
}

func NewCollidable(container *Object, mask CollidableMask) *Collidable {
	return &Collidable{
		Container: container,
		Mask:      mask,
	}
}

func (cr *Collidable) OnDraw(game *Game) error {
	if game.ShowHitboxes {
		cr.Mask.OnDraw(game)
	}
	return nil
}

func (rr *Collidable) OnUpdate(game *Game) error {
	rr.Mask.OnUpdate(game)
	return nil
}

func (rr *Collidable) AddCollisionListener(onColl func(obj *Object)) {
	rr.onCollision = append(rr.onCollision, onColl)
}
