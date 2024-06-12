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

type SceneCreator func(game *Game) Scene

type SceneObject struct {
	Scene        *Scene
	SceneCreator SceneCreator
}
