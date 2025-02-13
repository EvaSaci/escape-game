package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster struct {
	Name     string
	Position rl.Vector2
	Health   int
	Damage   int
	Loot     []item.Item
	Worth    int //valeur en argent quand tué
	Speed    float32

	IsAlive bool

	Sprite rl.Texture2D
}

func (m *Monster) Attack(p *Player) {
	p.Health -= m.Damage
}

func (m *Monster) ToString() {
	fmt.Printf("Je suis un monstre avec %d points de vie\n", m.Health)
}
