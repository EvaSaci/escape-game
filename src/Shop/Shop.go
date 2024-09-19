package shop

import (
	"fmt"
	"main/src/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)
func shop(player *entity.Player, s *entity.Shop) {
	if rl.IsKeyPressed(keyE)
		entity.Player.Money -= 50
		s.Inventory = append(s.Inventory, item.Item{Name: "Potion", Price: 25})
		entity.Player.Health += 25
}