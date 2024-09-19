package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
	WIN		 menu = iota
)

type engine int

const (
	INGAME  engine = iota
	PAUSE    engine = iota
	GAMEOVER engine = iota
)

type Engine struct {
	Player   entity.Player
	Monsters []entity.Monster
	Shopkeeper entity.Shop

	Music       rl.Music
	MusicVolume float32

	Sprites map[string]rl.Texture2D

	Camera rl.Camera2D

	MapJSON MapJSON

	LoadingScreenHome rl.Texture2D
	LoadingScreenPause rl.Texture2D

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}

const (
	ChaseDistance = 100 // distance between
)