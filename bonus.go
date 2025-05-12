package ebiten_pacwoman

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

type FruitType int

const (
	BANANA FruitType = iota
	APPLE
	CHERRY
)

type Bonus struct {
	ebiten_extended.Node2D
	sprite *ebiten_extended.Sprite
}

func NewBonus(fruitType FruitType) *Bonus {
	var sprite *ebiten_extended.Sprite
	switch fruitType {

	case BANANA:
		sprite = ebiten_extended.NewSprite("banana", ebiten_extended.ResourceManager().GetImage(BANANA_IMAGE), true)
	case APPLE:
		sprite = ebiten_extended.NewSprite("apple", ebiten_extended.ResourceManager().GetImage(APPLE_IMAGE), true)
	case CHERRY:
		sprite = ebiten_extended.NewSprite("cherry", ebiten_extended.ResourceManager().GetImage(CHERRY_IMAGE), true)
	}
	sprite.SetPivot(math2D.NewVector2D(15, 15))
	bonus := &Bonus { Node2D: *ebiten_extended.NewNode2D("bonus"), sprite: sprite }
	bonus.AddChild(sprite)
	return bonus
}
