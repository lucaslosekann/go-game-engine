package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Scene struct {
	name    string
	objects []*Object
	bgColor rl.Color
}

func NewScene(name string) Scene {
	return Scene{
		name: name,
	}
}
func (s *Scene) AddObject(o *Object) {
	s.objects = append(s.objects, o)
}

func (s *Scene) SetBgColor(c rl.Color) {
	s.bgColor = c
}

func (s *Scene) GetObjectByTag(tag string) *Object {
	for _, obj := range s.objects {
		if obj.Tag == tag {
			return obj
		}
	}
	return nil
}

type SceneEncapsulator interface {
	SceneCreator(game *Game) Scene
	OnLoad(game *Game, s *Scene)
	OnUnload(game *Game, s *Scene)
}

type SceneObject struct {
	Scene             *Scene
	SceneEncapsulator SceneEncapsulator
}
