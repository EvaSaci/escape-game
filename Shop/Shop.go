func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		if monster.Name == "Janne" {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 240, 420),
				rl.NewRectangle(monster.Position.X, monster.Position.Y ,240, 420),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)		
		} else {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 100, 100),
				rl.NewRectangle(monster.Position.X, monster.Position.Y, 150, 150),
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)
		}
		
	}
}
