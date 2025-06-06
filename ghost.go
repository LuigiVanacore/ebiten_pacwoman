package ebiten_pacwoman

import (
	"time"

	"github.com/LuigiVanacore/ebiten_extended"
)

type EnemyState int

const (
	STRONG EnemyState = iota
	WEAK
)

type Ghost struct {
	Character
	isWeak        bool
	weaknessTimer ebiten_extended.Timer
}

func NewGhost() *Ghost {
	animationPlayer := ebiten_extended.AnimationPlayer()
	movementComponent := NewMovementComponent()
	ghost := &Ghost{Character: NewCharacter("ghost", movementComponent, animationPlayer),
		isWeak:        false,
		weaknessTimer: ebiten_extended.NewTimer( time.Duration(0), false)}
}



// class Ghost : public Ch
// {
// 	public:

// 	enum State
// 	{
// 		Strong,
// 		Weak
// 	};

// 	public:
// 	Ghost(sf::Texture& texture, PacWoman* pacWoman);

// 	void setWeak(sf::Time duration);
// 	bool isWeak() const;

// 	void update(sf::Time delta);

// 	protected:
// 	void changeDirection();

// 	private:
// 	void draw(sf::RenderTarget& target, sf::RenderStates states) const;

// 	bool m_isWeak;
// 	sf::Time m_weaknessDuration;

// 	sf::Sprite m_visual;
// 	PacWoman* m_pacWoman;

// 	Animator m_strongAnimator;
// 	Animator m_weakAnimator;
// };
// #endif // PACWOMAN_GHOST_HPP



// Ghost::Ghost(sf::Texture& texture, PacWoman* pacWoman)
// :m_visual(texture)
// ,m_isWeak(false)
// ,m_weaknessDuration(sf::Time::Zero)
// ,m_pacWoman(pacWoman)

// {
// 	setOrigin(20, 20);
	
// 	//m_strongAnimator.addFrame(sf::IntRect(40, 32, 40, 40));
//     m_strongAnimator.addFrame(sf::IntRect(80, 32, 40, 40));

//     m_weakAnimator.addFrame(sf::IntRect(40, 72, 40, 40));
//     //m_weakAnimator.addFrame(sf::IntRect(80, 72, 40, 40));

//     m_strongAnimator.play(sf::seconds(0.25), true);
//     m_weakAnimator.play(sf::seconds(1), true);

// }

// void Ghost::setWeak(sf::Time duration)
// {
// 	m_isWeak =true;
// 	m_weaknessDuration = duration;
// }


// bool Ghost::isWeak() const
// {
// 	return m_isWeak;
// }

// void Ghost::draw(sf::RenderTarget& target, sf::RenderStates states) const
// {
// 	states.transform *= getTransform();
	
// 	target.draw(m_visual, states);
// }

// void Ghost::update(sf::Time delta)
// {    
	
// 	if (m_isWeak)
//     {
//         m_weaknessDuration -= delta;

//         if (m_weaknessDuration <= sf::Time::Zero)
//         {
//             m_isWeak = false;
//             m_strongAnimator.play(sf::seconds(0.25), true);
//         }
//     }
	
//     if (!m_isWeak)
//     {
//         m_strongAnimator.update(delta);
//         m_strongAnimator.animate(m_visual);
//     }
//     else
//     {
//         m_weakAnimator.update(delta);
//         m_weakAnimator.animate(m_visual);
//     }
    
//     Character::update(delta);
// }

// void Ghost::changeDirection() 
// {
// 	static sf::Vector2i directions[4] = {
//         sf::Vector2i(1, 0),
//         sf::Vector2i(0, 1),
//         sf::Vector2i(-1, 0),
//         sf::Vector2i(0, -1)
//         };
	
// 	std::map<float, sf::Vector2i> directionProb;
	
// 	float targetAngle;
	
// 	sf::Vector2f distance = m_pacWoman->getPosition() - getPosition();
	
// 	targetAngle = std::atan2(distance.x, distance.y) * (180/3.14);
	
// 	for (auto direction : directions)
// 	{
// 		float directionAngle = std::atan2(direction.x, direction.y) * (180/3.14);
		
// 		//Normalize the angle difference
// 		float diff = 180 - std::abs(std::abs(directionAngle - targetAngle) - 180);
		
// 		directionProb[diff] = direction;		
// 	}
// 	setDirection(directionProb.begin()->second);
	
// 	auto it = directionProb.begin();
	
// 	do
// 	{
// 		setDirection(it->second);
// 		it++;
// 	}
// 	while(!willMove());
// }
