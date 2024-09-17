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
		return
	} else if monster.Health <= 0 {
		player.Inventory = append(player.Inventory, monster.Loot...)
		player.Money += monster.Worth
		fmt.Println("--------------DEAD------------")
		monster.IsAlive = false
		monster.Health = 1
		return
	} else {
		if rl.IsKeyPressed(rl.KeyE) {
			player.Attack(monster)
			fmt.Println(monster.Health)
			//monster.ToString()   monster.Attack(player)
		}
			currentTime := float32(rl.GetTime())

			if currentTime - lastTime >= 1 {
				monster.Attack(player)
				fmt.Println(player.Health)
				lastTime = currentTime
			}
		}
}
