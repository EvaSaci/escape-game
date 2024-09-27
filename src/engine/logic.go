package engine

import (
	shop "main/src/Shop"
	"main/src/entity"
	"main/src/fight"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	//Menus
	if rl.IsKeyPressed(rl.KeyEnter) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyB) {
		e.StateMenu = HOME
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) InGameLogic() {
	// Mouvement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		e.Player.Position.Y -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		e.Player.Position.Y += e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		e.Player.Position.X -= e.Player.Speed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		e.Player.Position.X += e.Player.Speed
	}

	// Camera

	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}

	e.CheckCollisions()

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/extrait.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) CheckCollisions() {
	e.MonsterCollisions()
	e.UpdateMonsters()
	e.shopkeepercollision()
}
func (e *Engine) shopkeepercollision() {
	if e.Shopkeeper.Position.X > e.Player.Position.X-50 &&
		e.Shopkeeper.Position.X < e.Player.Position.X+50 &&
		e.Shopkeeper.Position.Y > e.Player.Position.Y-50 &&
		e.Shopkeeper.Position.Y < e.Player.Position.Y+50 {
		e.NormalTalke(e.Shopkeeper, "Je suis le shop du village!")
		shop.Shop(&e.Player, &e.Shopkeeper)
	} else {
	}
}

func (e *Engine) MonsterCollisions() {

	for i, monster := range e.Monsters {
		if monster.Worth > 0 {
			if monster.Position.X > e.Player.Position.X-40 &&
				monster.Position.X < e.Player.Position.X+40 &&
				monster.Position.Y > e.Player.Position.Y-40 &&
				monster.Position.Y < e.Player.Position.Y+40 {

				if monster.Health <= 0 {
					e.NormalTalk(monster, "Clique sur 'e' pour récupérer ta monnaie")
				} else {
					e.NormalTalk(monster, "Tu veut m'attaquer ?")
				}
				fight.Fight(&e.Player, &e.Monsters[i])

				if monster.Name == "mael" {
					e.RobotTalk(monster, "Tu veux m'attaquer?")
				}
				if monster.Name == "eva" {
					e.CipherTalk(monster, "Tu veux m'attaquer?")
				}

			}
		}
	}
}

func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)

}
func (e *Engine) NormalTalke(s entity.Shop, sentence string) {
	e.RenderDialoge(s, sentence)
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyA) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}
func (e *Engine) UpdateMonsters() {

	for i := 0; i < len(e.Monsters); i++ {

		if e.Monsters[i].IsAlive {
			distance := rl.Vector2Distance(e.Player.Position, e.Monsters[i].Position)

			if distance <= ChaseDistance {
				direction := rl.Vector2Subtract(e.Player.Position, e.Monsters[i].Position)
				direction = rl.Vector2Normalize(direction)
				e.Monsters[i].Position = rl.Vector2Add(e.Monsters[i].Position, direction)
			}
		}
	}
}
func (e *Engine) RobotTalk(m entity.Monster, sentence string) {
	var str2 string
	var l string
	for _, char := range sentence {
		x := int(char)
		l = ""
		for x/2 != 0 {
			r := x % 2
			l = strconv.Itoa(r) + l
			x = x / 2
		}
		l = "1" + l
		for len(l) < 8 {
			l = "0" + l
		}
		str2 = str2 + l
	}
	e.RenderDialog(m, str2)
}

func (e *Engine) CipherTalk(m entity.Monster, sentence string) {
	var str2 string
	for _, char := range sentence {
		char = char + 2
		str2 = str2 + string(char)
	}
	e.RenderDialog(m, str2)
}
