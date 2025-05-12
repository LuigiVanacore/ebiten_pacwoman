package ebiten_pacwoman

import (
    _ "embed"
)

var (
	//go:embed C:\Go_projects\ebiten_pacwoman\resources\logo.png
	logo []byte

//go:embed C:\Go_projects\ebiten_pacwoman\resources\pacwoman.png
pacwoman []byte

//go:embed C:\Go_projects\ebiten_pacwoman\resources\spritesheet.png
spriteSheet []byte

//go:embed C:\Go_projects\ebiten_pacwoman\resources\levels\large_level.png
large_level []byte


//go:embed C:\Go_projects\ebiten_pacwoman\resources\levels\large.png
large []byte


//go:embed C:\Go_projects\ebiten_pacwoman\resources\levels\level.png
level []byte


//go:embed C:\Go_projects\ebiten_pacwoman\resources\levels\medium.png
medium []byte

//go:embed C:\Go_projects\ebiten_pacwoman\resources\levels\small.png
small []byte
)