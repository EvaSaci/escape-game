package engine

import (
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Escape Raid")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 800, Y: 850},
		Health:    100,
		Money:     1000,
		Speed:     1.5,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "mael",
		Position: rl.Vector2{X: 130, Y: 800},
		Health:   50,
		Damage:   5,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/mael.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Eva",
		Position: rl.Vector2{X: 312, Y: 751},
		Health:   50,
		Damage:   5,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/eva.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "yasmine",
		Position: rl.Vector2{X: 407, Y: 607},
		Health:   50,
		Damage:   5,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/yas.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Janne",
		Position: rl.Vector2{X: 711, Y: 644},
		Health:   50,
		Damage:   5,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/janne.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Anderson",
		Position: rl.Vector2{X: 797, Y: 151},
		Health:   75,
		Damage:   5,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Anderson.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Yanisse",
		Position: rl.Vector2{X:	464, Y: 19},
		Health:   400,
		Damage:   40,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/yanisse.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Arnaud",
		Position: rl.Vector2{X: 435, Y: 331},
		Health:   75,
		Damage:   10,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/arnaud.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Léo",
		Position: rl.Vector2{X: 21, Y: 152},
		Health:   75,
		Damage:   10,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/leo.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Tom",
		Position: rl.Vector2{X: 642, Y: 482},
		Health:   75,
		Damage:   10,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/tom.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "anthony",
		Position: rl.Vector2{X: 50, Y: 301},
		Health:   75,
		Damage:   10,
		Speed:    0.9,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/antho.png"),
	})

	e.Shopkeeper =	entity.Shop{
		Name:      "Sébastien",
		Position:  rl.Vector2{X: 495, Y: 750},
		Inventory: []item.Item{},

		Sprite: rl.LoadTexture("textures/entities/shopkeeper/Shopkeeper.png"),
	}
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

	rl.PlayMusicStream(e.Music)
}
