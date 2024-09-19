package shop

import (
	"fmt"
	"main/src/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)
func shop(player *entity.Player, s *entity.Shop) {
	if rl.IsKeyPressed(rl.KeyE) {
		player.Health += 25
		player.Money -= 50
		fmt.Println("Tu vient de te heal 25 PV")
    } else {
		fmt.Println("Tu n'a pas l'argent!")

	}
}