package fight

import (
    "main/src/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

type fight int

const (
    PlayerTurn  fight = iota
    MonsterTurn fight = iota
)

func Fight(player entity.Player, monster *entity.Monster) {
	if player.Health <= 0 {
		player.IsAlive = false
		return
	}else if monster.Health <= 0 {
		player.Inventory = append(player.Inventory, monster.Loot...)
		player.Money += monster.Worth
		fmt.Println("--------------DEAD------------")
		monster.IsAlive = false
		monster.Health = 1
		return
	}else {
		if rl.IsKeyPressed(rl.KeyE) { 
			player.Attack(monster)
			fmt.Println(monster.Health)
			//monster.ToString()
		}
	}
		
}

