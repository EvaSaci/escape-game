package engine

import (
	"main/src/entity"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	BackGroudTexture rl.Texture2D
)

func (e *Engine) Rendering() {
	rl.ClearBackground(rl.Blue)
}

func (e *Engine) HomeRendering() {
	rl.DrawTexture(e.LoadingScreenHome, 0, 0, rl.White)
	rl.DrawText("Village Defend", int32(rl.GetScreenWidth())/2-rl.MeasureText("Village Defend", 40)/2-300, int32(rl.GetScreenHeight())/2-150, 60, rl.RayWhite)
	rl.DrawText("[Entrer] pour Jouer", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Enter] to Play", 20)/2-355, int32(rl.GetScreenHeight())/2, 30, rl.RayWhite)
	rl.DrawText("[Echap] pour Quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2-367, int32(rl.GetScreenHeight())/2+100, 30, rl.RayWhite)

	if rl.IsCursorOnScreen() {
		rl.HideCursor()
	}
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // On commence le rendu camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderShopkeeper()
	e.RenderPlayer()
	e.RenderHealth()
	rl.DrawFPS(int32(rl.GetScreenWidth())/2-700, int32(rl.GetScreenHeight())/2-500)

	if e.Player.Health <= 0 {
		e.GAMEOVER()
		return
	}
		for _, monster := range e.Monsters {
			if monster.Name == "Yanisse" {
				if monster.Health <= 0 {
					e.WIN()
					return
				}
			}
		}

	rl.EndMode2D() // On finit le rendu camera
	if e.Player.Health > 0 {
		// Ecriture fixe (car pas affectée par le mode camera)
		rl.DrawText("[Echap] Pause / Tutoriel", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to Pause", 20)/2+750, int32(rl.GetScreenHeight())/2-500, 22, rl.RayWhite)
		rl.DrawText("Monnaie :", int32(rl.GetScreenWidth())/2-rl.MeasureText("Monnaie :", 20)/2-750, int32(rl.GetScreenHeight())/2-350, 35, rl.RayWhite)
		rl.DrawText(strconv.Itoa(e.Player.Money), int32(rl.GetScreenWidth())/2+rl.MeasureText("Monnaie :", 20)/2-670, int32(rl.GetScreenHeight())/2-350, 35, rl.RayWhite)
	}
}

func (e *Engine) PauseRendering() {
	rl.DrawTexture(e.LoadingScreenPause, 0, 0, rl.White)

	rl.DrawText("Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("Paused", 40)/2-300, int32(rl.GetScreenHeight())/2-150, 40, rl.White)
	rl.DrawText("[Echap] pour revenir en jeux", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] or [Esc] to resume", 20)/2-255, int32(rl.GetScreenHeight())/2, 20, rl.White)
	rl.DrawText("[Q] pour quitter", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 20)/2-302, int32(rl.GetScreenHeight())/2+100, 20, rl.White)
}

func (e *Engine) GAMEOVER() {
	rl.DrawTexture(e.LoadingScreenPause, 0, 0, rl.White)

	rl.DrawText("GAME OVER", int32(rl.GetScreenWidth())/2-rl.MeasureText("GAME OVER", 40)/2, int32(rl.GetScreenHeight())/2-150, 80, rl.RayWhite)
	rl.DrawText("Appuis sur [Echap] + [A]/[Q]", int32(rl.GetScreenWidth())/2-rl.MeasureText("Appuis sur [Echap] + [A]/[Q]", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.White)
}

func (e *Engine) WIN() {
	rl.ClearBackground(rl.Green)

    rl.DrawText("Vous avez gagné!", int32(rl.GetScreenWidth())/2-rl.MeasureText("Vous avez gagné!", 40)/2, int32(rl.GetScreenHeight())/2-150, 80, rl.RayWhite)
    rl.DrawText("Appuis sur [Echap] + [A]/[Q]", int32(rl.GetScreenWidth())/2-rl.MeasureText("Appuis sur [Echap] + [A]/[Q]", 20)/2, int32(rl.GetScreenHeight())/2+100, 20, rl.White)
}
func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite,
		rl.NewRectangle(0, 0, 100, 100),
		rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		if monster.Worth > 0 {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 100, 100),
				rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)
		}
	}
}


func (e *Engine) RenderShopkeeper() {
	rl.DrawTexturePro(
		e.Shopkeeper.Sprite,
		rl.NewRectangle(0, 0, 16, 28),
		rl.NewRectangle(e.Shopkeeper.Position.X, e.Shopkeeper.Position.Y, 32, 56),
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)
}

func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
}

func (e *Engine) RenderDialoge(s entity.Shop, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(s.Position.X)-20,
		int32(s.Position.Y),
		10,
		rl.RayWhite,
	)
	rl.EndMode2D()
}

func (e *Engine) RenderHealth() {
	if e.Player.Health < 0 {
		e.Player.Health = 0
	} else if e.Player.Health > 100 {
		e.Player.Health = 100
	}

	rl.DrawRectangle(int32(e.Player.Position.X)+25, int32(e.Player.Position.Y)+30, int32(100), 5, rl.DarkBrown)
	rl.DrawText(strconv.Itoa(e.Player.Health), int32(e.Player.Position.X)+25, int32(e.Player.Position.Y)+35, int32(3), rl.Pink)
	rl.DrawRectangle(int32(e.Player.Position.X)+25, int32(e.Player.Position.Y)+30, int32(e.Player.Health), 5, rl.Red)
	for _, monster := range e.Monsters {
		if monster.IsAlive {
			if monster.Health < 0 {
				monster.Health = 0
			} else if monster.Health > 400 {
				monster.Health = 400
			}
			

			rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(50), 5, rl.DarkBrown)
			rl.DrawText(strconv.Itoa(monster.Health), int32(monster.Position.X)+25, int32(monster.Position.Y)+35, int32(3), rl.White)
			rl.DrawRectangle(int32(monster.Position.X)+25, int32(monster.Position.Y)+30, int32(monster.Health), 5, rl.Red)
		}
	}
	rl.EndMode2D()
}
