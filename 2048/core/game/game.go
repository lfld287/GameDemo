package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"tomatoSister/2048/core/game/board"
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
	for i := 0; i < len(keyList); i++ {
		fmt.Printf("key: %v\n", keyList[i])
		p.board.AppendInput(keyList[i])
	}
	return nil
}

func (p *Game) Draw(screen *ebiten.Image) {
	p.board.Draw(screen)
}

func (p *Game) Layout(outW, outH int) (screenWidth, screenHeight int) {
	return outW, outH
}
