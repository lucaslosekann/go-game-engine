package engine

import (
	"fmt"
	"reflect"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Object struct {
	Components []Component
	Active     bool
	Position   rl.Vector2
	Anchor     rl.Vector2
	Rotation   float32
	Scale      float32
	Collidable *Collidable
	Tag        string
}

func (o *Object) Update(game *Game) error {
	//On Update reverse order to update the last object added first

	for i := len(o.Components) - 1; i >= 0; i-- {
		comp := o.Components[i]
		if err := comp.OnUpdate(game); err != nil {
			return err
		}
	}

	// Update RigidBody
	if o.Collidable != nil {
		if err := o.Collidable.OnUpdate(game); err != nil {
			return err
		}
	}

	return nil
}

func (o *Object) Draw(game *Game) error {
	for _, comp := range o.Components {
		err := comp.OnDraw(game)
		if err != nil {
			return err
		}
	}

	if o.Collidable != nil {
		if err := o.Collidable.OnDraw(game); err != nil {
			return err
		}
	}

	return nil
}

func (o *Object) AddComponent(new Component) {
	typ := reflect.TypeOf(new)
	for _, existing := range o.Components {
		if reflect.TypeOf(existing) == typ {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	o.Components = append(o.Components, new)
}

func (o *Object) GetComponent(dummy Component) Component {
	typ := reflect.TypeOf(dummy)
	for _, existing := range o.Components {
		if reflect.TypeOf(existing) == typ {
			return existing
		}
	}
	panic(fmt.Sprintf("no component with type %v", typ))
}

func (o *Object) HasComponent(dummy Component) bool {
	typ := reflect.TypeOf(dummy)
	for _, existing := range o.Components {
		if reflect.TypeOf(existing) == typ {
			return true
		}
	}
	return false
}
