package block

import "github.com/hajimehoshi/ebiten/v2"

type blockState interface {
	update() (nextState blockState, updating bool)
	draw(screen *ebiten.Image) (image *ebiten.Image)
}
