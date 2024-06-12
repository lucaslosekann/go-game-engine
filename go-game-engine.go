package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	DrawRectangle = rl.DrawRectangle
)

type Game struct {
	running      bool
	Title        string
	ShowHitboxes bool
	activeScene  *Scene
	scenes       map[string]*SceneObject
}

func NewGame(title string) *Game {
	return &Game{
		running: false,
		Title:   title,
		scenes:  map[string]*SceneObject{},
	}
}

func (g *Game) GetActiveScene() *Scene {
	return g.activeScene
}

func (g *Game) SetActiveScene(scene string) error {
	s, ok := g.scenes[scene]
	if !ok {
		return fmt.Errorf("scene not loaded")
	}
	g.activeScene = s.Scene
	return nil
}

func (g *Game) Init() error {
	if !rl.IsWindowReady() {
		return fmt.Errorf("window not initialized")
	}

	if g.activeScene == nil {
		return fmt.Errorf("no initial scene")
	}
	g.running = true

	for !rl.WindowShouldClose() && g.running {

		rl.BeginDrawing()
		rl.ClearBackground(g.activeScene.bgColor)

		for i, object := range g.activeScene.objects {
			if object.Active {
				if err := object.Update(g); err != nil {
					return err
				}

				// Check collisions
				for j := i + 1; j < len(g.activeScene.objects); j++ {
					if g.activeScene.objects[j].Active {
						if object.Collidable != nil {
							if object.Collidable.Mask.CheckCollision(g.activeScene.objects[j].Collidable.Mask) {
								for _, onColl := range object.Collidable.onCollision {
									onColl(g.activeScene.objects[j])
								}
								for _, onColl := range g.activeScene.objects[j].Collidable.onCollision {
									onColl(object)
								}
							}
						}
					}
				}

				if err := object.Draw(g); err != nil {
					return err
				}
			}
		}

		rl.EndDrawing()
	}
	return nil
}

func (g *Game) AddScene(creator SceneCreator) {
	s := creator(g)
	g.scenes[s.name] = &SceneObject{
		SceneCreator: creator,
		Scene:        &s,
	}
}

func (g *Game) ReloadScene(scene string) error {
	s, ok := g.scenes[scene]
	if !ok {
		return fmt.Errorf("scene not loaded")
	}

	*s.Scene = s.SceneCreator(g)
	return nil
}

func (g *Game) ReloadCurrentScene() {
	g.ReloadScene(g.activeScene.name)
}
