package ebiten_pacwoman

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)


type CellData int

const (
	EMPTY CellData = iota
	WALL
	DOT
	SUPERDOT
	BONUS
)


type Maze struct {
	ebiten_extended.Node2D
	size math2D.Vector2D
	mazeData []CellData
	pacWomanPosition math2D.Vector2D
	ghostPositions []math2D.Vector2D
	sprite *ebiten_extended.SpriteNode
}


func NewMaze(size math2D.Vector2D, maze_sprite *ebiten.Image) *Maze {
	sprite := ebiten_extended.NewSprite("maze_sprite", maze_sprite, true)
	maze := &Maze{
		Node2D: *ebiten_extended.NewNode2D("maze"),
		size: size,
		mazeData: make([]CellData, int(size.X()*size.Y())),
		pacWomanPosition: math2D.NewVector2D(0, 0),
		ghostPositions: []math2D.Vector2D{},
		sprite: sprite,
	}
	maze.AddChild(sprite)
	return maze
}

func (m *Maze) LoadLevel(name string) {
	// Load the level data from a file or other source
	// For now, we will just initialize some dummy data
	for i := 0; i < len(m.mazeData); i++ {
		m.mazeData[i] = EMPTY
	}
	m.pacWomanPosition = math2D.NewVector2D(1, 1)
	m.ghostPositions = []math2D.Vector2D{
		math2D.NewVector2D(5, 5),
		math2D.NewVector2D(6, 6),
	}
}

func (m *Maze) GetPacWomanPosition() math2D.Vector2D {
	return m.pacWomanPosition
}

func (m *Maze) GetGhostPositions() []math2D.Vector2D {
	return m.ghostPositions
}

func (m *Maze) PositionToIndex(position math2D.Vector2D) int {
	index := int(position.Y() * m.size.X() + position.X())
	if index < 0 || index >= len(m.mazeData) {
		return -1 // Invalid index
	}
	return index
}

func (m *Maze) IndexToPosition(index int) math2D.Vector2D {
	if index < 0 || index >= len(m.mazeData) {
		return math2D.NewVector2D(-1, -1) // Invalid position
	}
	x := index % int(m.size.X())
	y := index / int(m.size.X())
	return math2D.NewVector2D(float64(x), float64(y))
}


func (m *Maze) MapPixelToCell(pixel math2D.Vector2D) math2D.Vector2D {
	return math2D.NewVector2D(pixel.X() / 32, pixel.Y() / 32)
}

func (m *Maze) MapCellToPixel(cell math2D.Vector2D) math2D.Vector2D {
	return math2D.NewVector2D(cell.X() * 32, cell.Y() * 32)
}

func (m *Maze) IsWall(position math2D.Vector2D) bool {
	index := m.PositionToIndex(position)
	if index < 0 {
		return false // Invalid position
	}
	return m.mazeData[index] == WALL
}

func (m *Maze) IsDot(position math2D.Vector2D) bool {
	index := m.PositionToIndex(position)
	if index < 0 {
		return false // Invalid position
	}
	return m.mazeData[index] == DOT
}

func (m *Maze) IsSuperDot(position math2D.Vector2D) bool {
	index := m.PositionToIndex(position)
	if index < 0 {
		return false // Invalid position
	}
	return m.mazeData[index] == SUPERDOT
}

func (m *Maze) PickObject(position math2D.Vector2D) {
	index := m.PositionToIndex(position)
	if index < 0 {
		return // Invalid position
	}
	if m.mazeData[index] == DOT {
		m.mazeData[index] = EMPTY
	} else if m.mazeData[index] == SUPERDOT {
		m.mazeData[index] = EMPTY
	}
}


func (m *Maze) IsBonus(position math2D.Vector2D) bool {
	index := m.PositionToIndex(position)
	if index < 0 {
		return false // Invalid position
	}
	return m.mazeData[index] == BONUS
}

func (m *Maze) GetSize() math2D.Vector2D {
	return m.size
}

func (m *Maze) GetRemainingDots() int {
	count := 0
	for _, cell := range m.mazeData {
		if cell == DOT {
			count++
		}
	}
	return count
}



// class Maze : public sf::Drawable
// {
// 	public:
	
// 	Maze(sf::Texture& texture);
// 	void loadLevel(std::string name);
	
// 	sf::Vector2i getPacWomanPosition() const;
// 	std::vector<sf::Vector2i> getGhostPositions() const;
	
// 	inline std::size_t positionToIndex(sf::Vector2i position) const;
// 	inline sf::Vector2i indexToPosition(std::size_t index) const;
	
// 	sf::Vector2i mapPixelToCell(sf::Vector2f pixel) const;
// 	sf::Vector2f mapCellToPixel(sf::Vector2i cell) const;
	
// 	bool isWall(sf::Vector2i position) const;
// 	bool isDot(sf::Vector2i position) const;
// 	bool isSuperDot(sf::Vector2i position) const;
// 	void pickObject(sf::Vector2i position);
	
// 	bool isBonus(sf::Vector2i position) const;
	
// 	sf::Vector2i getSize() const;
	
// 	int getRemainingDots() const;
		
// 	private:
	
// 	enum CellData
// 	{
// 		Empty,
// 		Wall,
// 		Dot,
// 		SuperDot,
// 		Bonus
// 	};
	
// 	void draw(sf::RenderTarget& target, sf::RenderStates states) const;
	
// 	sf::Vector2i m_mazeSize;
// 	std::vector<CellData> m_mazeData;
// 	sf::Vector2i m_pacWomanPosition;
// 	std::vector<sf::Vector2i> m_ghostPositions;
		
// 	sf::RenderTexture m_renderTexture;
// 	sf::Texture& m_texture;
// };

// #endif // PACWOMAN_MAZE_HPP
