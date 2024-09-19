package entity

import (
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shop struct {

	Name     string
	Position  rl.Vector2
	Money     int
	Inventory []item.Item

	Sprite rl.Texture2D
}

func (s *Shop) Buy(item item.Item, player *Player) {
	shop := s
	if shop.Money >= item.Price {
        shop.Money -= item.Price
        player.Money += item.Price
        player.Inventory = append(player.Inventory, item)
    }
}

func (s *Shop) shopkeeper(p *Player) {
	p.Health += 25
	p.Money -= 50
}