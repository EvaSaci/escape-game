package fight

import (
	"fmt"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type fight int

const (
	PlayerTurn  fight = iota
	MonsterTurn fight = iota
)

var lastTime float32 = 0

func Fight(player *entity.Player, monster *entity.Monster) {
	if player.Health <= 0 {
		player.IsAlive = false
		player.Health = 0
		player.Sprite = rl.LoadTexture("textures/entities/soldier/Soldier-Death.png")
		return

	} else if monster.Health <= 0 {
		player.Inventory = append(player.Inventory, monster.Loot...)
		player.Money += monster.Worth
		monster.IsAlive = false
		monster.Health = 0
		monster.Sprite = rl.LoadTexture("textures/entities/orc/Orc-Death.png")
		return
	} else {
		if rl.IsKeyPressed(rl.KeyE) {
			player.Attack(monster)
			monster.Sprite = rl.LoadTexture("textures/entities/orc/Orc-Idle.png")
		}
			currentTime := float32(rl.GetTime())

			if currentTime - lastTime >= 1 {
				monster.Attack(player)
				fmt.Println(player.Health)
				lastTime = currentTime
			}
		}
}
