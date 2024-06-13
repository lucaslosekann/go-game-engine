package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	running      bool
	Title        string
	ShowHitboxes bool
	activeScene  *Scene
	scenes       map[string]*SceneObject
	textures     map[string]*rl.Texture2D
}

func NewGame(title string) *Game {
	return &Game{
		running: false,
		Title:   title,
		scenes:  map[string]*SceneObject{},
	}
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
				if object.Collidable != nil {
					for j := i + 1; j < len(g.activeScene.objects); j++ {
						if g.activeScene.objects[j].Active && g.activeScene.objects[j].Collidable != nil {
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

// Scene
func (g *Game) GetScene(scene string) (*Scene, bool) {
	s, ok := g.scenes[scene]
	if !ok {
		return nil, false
	}
	return s.Scene, true
}
func (g *Game) GetActiveScene() *Scene {
	return g.activeScene
}
func (g *Game) SetActiveScene(scene string) error {
	s, ok := g.scenes[scene]
	if !ok {
		return fmt.Errorf("scene not loaded")
	}
	// Unload current scene
	if g.activeScene != nil {
		active := g.scenes[g.activeScene.name]
		active.SceneEncapsulator.OnUnload(g, active.Scene)
	}
	// Load new scene
	g.activeScene = s.Scene

	rl.SetMouseCursor(rl.MouseCursorDefault)
	s.SceneEncapsulator.OnLoad(g, s.Scene)
	return nil
}
func (g *Game) AddScene(encapsulator SceneEncapsulator) {
	s := encapsulator.SceneCreator(g)
	g.scenes[s.name] = &SceneObject{
		SceneEncapsulator: encapsulator,
		Scene:             &s,
	}
}
func (g *Game) RecreateScene(scene string) error {
	s, ok := g.scenes[scene]
	if !ok {
		return fmt.Errorf("scene not loaded")
	}

	*s.Scene = s.SceneEncapsulator.SceneCreator(g)
	return nil
}
func (g *Game) ReloadCurrentScene() {
	g.RecreateScene(g.activeScene.name)
}

func (g *Game) GetTexture(path string) *rl.Texture2D {
	if _, ok := g.textures[path]; !ok {
		t := rl.LoadTexture(path)
		g.textures[path] = &t
	}

	return g.textures[path]
}
