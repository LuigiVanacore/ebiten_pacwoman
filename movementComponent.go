package ebiten_pacwoman

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"math"
)

type MovementComponent struct {
	speed                float64
	previousIntersection math2D.Vector2D
	availableDirections  [4]bool
	currentDirection     math2D.Vector2D
	nextDirection        math2D.Vector2D
	hitBox               math2D.Rectangle
	willMove             bool
}

func NewMovementComponent(speed float64, hitBox math2D.Rectangle) MovementComponent {
	return MovementComponent{
		speed:                speed,
		currentDirection:     math2D.NewVector2D(1, 0),
		nextDirection:        math2D.NewVector2D(0, 0),
		previousIntersection: math2D.NewVector2D(0, 0),
		hitBox:               hitBox,
	}
}

func (c *MovementComponent) SetSpeed(speed float64) {
	c.speed = speed
}

func (c *MovementComponent) GetSpeed() float64 {
	return c.speed
}

func (c *MovementComponent) SetDirection(direction math2D.Vector2D) {
	c.nextDirection = direction
}

func (c *MovementComponent) GetDirection() math2D.Vector2D {
	return c.currentDirection
}

func (c *MovementComponent) SetWillMove(willMove bool) {
	c.willMove = willMove
}

func (c *MovementComponent) WillMove() bool {
	return c.willMove
}

func (c *MovementComponent) UpdateMovement(maze *Maze) {
	pixelPosition := c.GetPosition()

	pixelTraveled := c.GetSpeed()

	nextPixelPosition := math2D.AddVectors(pixelPosition, c.currentDirection.MultiplyScalar(pixelTraveled))
	c.SetPosition(nextPixelPosition)

	cellPosition := maze.MapPixelToCell(pixelPosition)

	offset := math2D.NewVector2D(
		math.Mod(pixelPosition.X(), 32)-16,
		math.Mod(pixelPosition.Y(), 32)-16,
	)

	if maze.IsWall(math2D.AddVectors(cellPosition, c.currentDirection)) {
		if (c.currentDirection.X() == 1 && offset.X() > 0) ||
			(c.currentDirection.X() == -1 && offset.X() < 0) ||
			(c.currentDirection.Y() == 1 && offset.Y() > 0) ||
			(c.currentDirection.Y() == -1 && offset.Y() < 0) {

			c.SetPosition(maze.MapCellToPixel(cellPosition))
		}
	}

	if !maze.IsWall(math2D.AddVectors(cellPosition, c.nextDirection)) && c.currentDirection != c.nextDirection {
		if (c.currentDirection.Y() == 0 && (offset.X() > -2 && offset.X() < 2)) ||
			(c.currentDirection.X() == 0 && (offset.Y() > -2 && offset.Y() < 2)) {
			c.SetPosition(maze.MapCellToPixel(cellPosition))
			c.currentDirection = c.nextDirection
			c.UpdateRotationAndScale()
		}
	}

	directions := [4]math2D.Vector2D{
		math2D.NewVector2D(1, 0),
		math2D.NewVector2D(0, 1),
		math2D.NewVector2D(-1, 0),
		math2D.NewVector2D(0, -1),
	}

	if cellPosition != c.previousIntersection {
		if (c.currentDirection.Y() == 0 && (offset.X() > -2 && offset.X() < 2)) ||
			(c.currentDirection.X() == 0 && (offset.Y() > -2 && offset.Y() < 2)) {
			var availableDirections [4]bool

			for i, direction := range directions {
				availableDirections[i] = maze.IsWall(math2D.AddVectors(cellPosition, direction))
			}

			if c.availableDirections != availableDirections {
				c.previousIntersection = cellPosition
				c.availableDirections = availableDirections
				c.ChangeDirection()
			}
		}
	}
}

func (c *MovementComponent) UpdateRotationAndScale() {
	if c.currentDirection.IsEqual(math2D.NewVector2D(1, 0)) {
		c.SetRotation(0)
		c.SetScale(-1, 1)
	} else if c.currentDirection.IsEqual(math2D.NewVector2D(0, 1)) {
		c.SetRotation(90)
		c.SetScale(-1, 1)
	} else if c.currentDirection.IsEqual(math2D.NewVector2D(-1, 0)) {
		c.SetRotation(0)
		c.SetScale(1, 1)
	} else if c.currentDirection.IsEqual(math2D.NewVector2D(0, -1)) {
		c.SetRotation(90)
		c.SetScale(1, 1)
	}
}

func (m *MovementComponent) ChangeDirection() {
	if m.availableDirections[0] && m.currentDirection != math2D.NewVector2D(1, 0) {
		m.nextDirection = math2D.NewVector2D(1, 0)
	} else if m.availableDirections[1] && m.currentDirection != math2D.NewVector2D(0, 1) {
		m.nextDirection = math2D.NewVector2D(0, 1)
	} else if m.availableDirections[2] && m.currentDirection != math2D.NewVector2D(-1, 0) {
		m.nextDirection = math2D.NewVector2D(-1, 0)
	} else if m.availableDirections[3] && m.currentDirection != math2D.NewVector2D(0, -1) {
		m.nextDirection = math2D.NewVector2D(0, -1)
	}
}

func (c *MovementComponent) GetHitbox() math2D.Rectangle {
	return c.hitBox
}
