package board

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"image/color"
)

var backgroundRGBA = color.RGBA{R: 127, G: 127, B: 127, A: 0xff}

var textColor = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

var blockColor = color.RGBA{R: 150, A: 0xff}

type Board struct {
	blockTextFont  font.Face
	blockTextColor color.Color

	blocks [4][4]*Block
}

func NewBoard() *Board {
	res := &Board{
		blockTextFont:  bigFont,
		blockTextColor: color.White,
		blocks:         [4][4]*Block{},
	}
	res.generateRandomBlockAtEmptySlot()
	return res
}

func (p *Board) Draw(screen *ebiten.Image) {
	p.drawBackground(screen)
	p.drawBlocks(screen)
}

func (p *Board) drawBackground(screen *ebiten.Image) {
	screen.Fill(backgroundRGBA)
}

func (p *Board) drawBlocks(screen *ebiten.Image) {
	p.drawBackground(screen)
	//split screen into 16 blocks
	//draw each block
	geoM := ebiten.GeoM{}
	widthUnit := float64(screen.Bounds().Size().X) / 4
	heightUnit := float64(screen.Bounds().Size().Y) / 4
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if p.blocks[row][col] == nil {
				continue
			}
			img := p.blocks[row][col].Render()
			geoM.Reset()
			geoM.Scale(widthUnit/100, heightUnit/100)
			geoM.Translate(float64(col)*widthUnit, float64(row)*heightUnit)
			screen.DrawImage(img, &ebiten.DrawImageOptions{
				GeoM: geoM,
			})
		}
	}
}

func (p *Board) generateRandomBlockAtEmptySlot() {
	c := 0
	var emptySlot [16]struct {
		row int
		col int
	}

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if p.blocks[row][col] == nil {
				emptySlot[c].row = row
				emptySlot[c].col = col
				c += 1
			}
		}
	}

	if c == 0 {
		return
	}

	r := generateRandomNumber(c)

	p.blocks[emptySlot[r].row][emptySlot[r].col] = NewBlock(2, blockColor, p.blockTextFont, p.blockTextColor)
}

func (p *Board) AppendInput(key ebiten.Key) {
	switch key {
	case ebiten.KeyDown:
		p.processVertical(true)
	case ebiten.KeyUp:
		p.processVertical(false)
	case ebiten.KeyRight:
		p.processHorizontal(true)
	case ebiten.KeyLeft:
		p.processHorizontal(false)
	}
}

func (p *Board) processVertical(down bool) {
	changed := false

	for col := 0; col < len(p.blocks[0]); col++ {
		original := make([]int, len(p.blocks))
		for row := 0; row < len(original); row++ {
			if p.blocks[row][col] == nil {
				original[row] = 0
				continue
			}
			original[row] = p.blocks[row][col].number
		}

		processed := processLine(original, down)

		fmt.Printf("col(%d) from : %v to %v\n", col, original, processed)

		for row := 0; row < len(processed); row++ {
			if processed[row] == original[row] {
				continue
			}
			changed = true
			if processed[row] == 0 {
				p.blocks[row][col] = nil
			} else {
				p.blocks[row][col] = NewBlock(processed[row], blockColor, p.blockTextFont, p.blockTextColor)
			}
		}
	}

	if changed {
		p.generateRandomBlockAtEmptySlot()
	}
}

func (p *Board) processHorizontal(right bool) {

	changed := false
	for row := 0; row < len(p.blocks); row++ {
		original := make([]int, len(p.blocks[row]))
		for col := 0; col < len(original); col++ {
			if p.blocks[row][col] == nil {
				original[col] = 0
				continue
			}
			original[col] = p.blocks[row][col].number
		}

		processed := processLine(original, right)
		fmt.Printf("row(%d) from : %v to %v\n", row, original, processed)

		for col := 0; col < len(processed); col++ {
			if processed[col] == original[col] {
				continue
			}
			changed = true
			if processed[col] == 0 {
				p.blocks[row][col] = nil
			} else {
				p.blocks[row][col] = NewBlock(processed[col], blockColor, p.blockTextFont, p.blockTextColor)
			}
		}
	}

	if changed {
		p.generateRandomBlockAtEmptySlot()
	}
}
