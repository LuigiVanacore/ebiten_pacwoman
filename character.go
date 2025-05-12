package ebiten_pacwoman

import (
	"github.com/LuigiVanacore/ebiten_extended"
)

type Character struct {
	ebiten_extended.Node2D
	animationPlayer *ebiten_extended.AnimationPlayer
	movementComponent MovementComponent
}


func NewCharacter(name string, movementComponent MovementComponent, animationPlayer *ebiten_extended.AnimationPlayer) Character {
	return Character{
		Node2D: *ebiten_extended.NewNode2D(name),
		movementComponent: movementComponent,
		animationPlayer: animationPlayer,
	}
}