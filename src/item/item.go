package item

import (
	"fmt"
)

type Item struct {
	Name         string
	Price        int
	IsConsumable bool
	IsEquippable bool
}

func (i *Item) Potion() {
	fmt.Printf("%s est utilisé pour soigner\n", i.Name,)
}

