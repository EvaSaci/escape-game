package shop

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Shop(player *entity.Player, s *entity.Shop) {
	if rl.IsKeyPressed(rl.KeyE) {
		if player.Money >= 50 {
			player.Money -= 50
			player.Health += 25
		}
	}
}
