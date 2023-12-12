package game

import (
	"GameDemo/2048/core/game/board"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	board *board.Board
}

func NewGame() *Game {
	return &Game{
		board: board.NewBoard(),
	}
}

func (p *Game) Update() error {
	var keyList = make([]ebiten.Key, 0, 10)
	keyList = inpututil.AppendJustPressedKeys(keyList)

	p.board.Update(keyList)

	return nil
}

func (p *Game) Draw(screen *ebiten.Image) {
	p.board.Draw(screen)
}

func (p *Game) Layout(outW, outH int) (screenWidth, screenHeight int) {
	smaller := outW
	if smaller > outH {
		smaller = outH
	}
	return smaller, smaller
}
