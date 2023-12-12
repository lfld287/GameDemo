package board

import (
	"GameDemo/2048/core/game/board/block"
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

	blocks  [4][4]*block.Block
	changed bool

	updating bool
}

func NewBoard() *Board {
	res := &Board{
		blockTextFont:  bigFont,
		blockTextColor: color.White,
		blocks:         [4][4]*block.Block{},
		changed:        false,
		updating:       false,
	}
	res.changed = true
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
	//TODO use row col
	widthUnit := float64(screen.Bounds().Size().X) / 4
	heightUnit := float64(screen.Bounds().Size().Y) / 4
	for row := 0; row < len(p.blocks); row++ {
		for col := 0; col < len(p.blocks[row]); col++ {
			if p.blocks[row][col] == nil {
				continue
			}
			p.blocks[row][col].Draw(screen, widthUnit, heightUnit)
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

	p.blocks[emptySlot[r].row][emptySlot[r].col] = block.NewBlock(2,
		block.LogicalPosition{
			X: emptySlot[r].col,
			Y: emptySlot[r].row,
		}, blockColor, p.blockTextFont, p.blockTextColor)
}

func (p *Board) Update(keyList []ebiten.Key) {
	p.updating = false
	for row := 0; row < len(p.blocks); row++ {
		for col := 0; col < len(p.blocks[row]); col++ {
			if p.blocks[row][col] == nil {
				continue
			}
			if !p.blocks[row][col].Update() {
				p.updating = true
			}
		}
	}

	if p.updating {
		return
	}

	if p.changed {
		p.changed = false
		newBlocks := [4][4]*block.Block{}
		for row := 0; row < len(p.blocks); row++ {
			for col := 0; col < len(p.blocks[row]); col++ {
				if p.blocks[row][col] == nil {
					continue
				}
				b := p.blocks[row][col]
				pos := b.GetPos()
				val := b.GetValue()
				if newBlocks[pos.Y][pos.X] != nil {
					newBlocks[pos.Y][pos.X].Show(val + newBlocks[pos.Y][pos.X].GetValue())
				} else {
					newBlocks[pos.Y][pos.X] = b
				}
			}
		}

		p.blocks = newBlocks
		p.generateRandomBlockAtEmptySlot()
		return
	}

	for i := 0; i < len(keyList); i++ {
		switch keyList[i] {
		case ebiten.KeyDown:
			p.changed = p.moveVertical(true)
			break
		case ebiten.KeyUp:
			p.changed = p.moveVertical(false)
			break
		case ebiten.KeyRight:
			p.changed = p.moveHorizontal(true)
			break
		case ebiten.KeyLeft:
			p.changed = p.moveHorizontal(false)
			break
		default:
			continue
		}
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
			original[row] = p.blocks[row][col].Number
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
				p.blocks[row][col] = block.NewBlock(processed[row], block.LogicalPosition{
					X: col,
					Y: row,
				}, blockColor, p.blockTextFont, p.blockTextColor)
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
			original[col] = p.blocks[row][col].Number
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
				p.blocks[row][col] = block.NewBlock(processed[col], block.LogicalPosition{
					X: col,
					Y: row,
				}, blockColor, p.blockTextFont, p.blockTextColor)
			}
		}
	}

	if changed {
		p.generateRandomBlockAtEmptySlot()
	}
}

func (p *Board) moveVertical(down bool) bool {
	moved := false
	for col := 0; col < len(p.blocks[0]); col++ {
		line := make([]*block.Block, len(p.blocks))
		for row := 0; row < len(line); row++ {
			line[row] = p.blocks[row][col]
		}

		if down {
			if p.moveLineVertical(line) {
				moved = true
			}
		} else {
			if p.moveLineVerticalReverse(line) {
				moved = true
			}
		}
	}
	return moved
}

func (p *Board) moveHorizontal(right bool) bool {
	moved := false
	for row := 0; row < len(p.blocks); row++ {
		line := make([]*block.Block, len(p.blocks[row]))
		for col := 0; col < len(line); col++ {
			line[col] = p.blocks[row][col]
		}

		if right {
			if p.moveLineHorizontal(line) {
				moved = true
			}
		} else {
			if p.moveLineHorizontalReverse(line) {
				moved = true
			}
		}
	}
	return moved
}

func (p *Board) moveLineVertical(line []*block.Block) bool {
	currentPos := len(line) - 1
	currentPosStack := 0
	currentPosValue := 0

	moved := false

	insertBlock := func(b *block.Block) {
		if b == nil {
			return
		}

		if currentPosStack == 0 {
			currentPosValue += b.Number
			currentPosStack += 1
			if b.MoveVertical(currentPos) {
				moved = true
			}
			return
		}

		if currentPosStack == 1 && currentPosValue == b.GetValue() {

			currentPosValue += b.Number
			currentPosStack += 1
			if b.MoveVertical(currentPos) {
				moved = true
			}
			return

		}

		currentPos -= 1
		currentPosValue = b.Number
		currentPosStack = 1
		if b.MoveVertical(currentPos) {
			moved = true
		}
		return

	}

	for i := len(line) - 1; i >= 0; i-- {
		insertBlock(line[i])
	}

	return moved
}

func (p *Board) moveLineVerticalReverse(line []*block.Block) bool {
	currentPos := 0
	currentPosStack := 0
	currentPosValue := 0

	moved := false

	insertBlock := func(b *block.Block) {
		if b == nil {
			return
		}

		if currentPosStack == 0 {
			currentPosValue += b.Number
			currentPosStack += 1
			if b.MoveVertical(currentPos) {
				moved = true
			}
			return
		}

		if currentPosStack == 1 && currentPosValue == b.GetValue() {

			currentPosValue += b.Number
			currentPosStack += 1
			if b.MoveVertical(currentPos) {
				moved = true
			}
			return

		}

		currentPos += 1
		currentPosValue = b.Number
		currentPosStack = 1
		if b.MoveVertical(currentPos) {
			moved = true
		}
		return

	}

	for i := 0; i < len(line); i++ {
		insertBlock(line[i])
	}

	return moved
}

func (p *Board) moveLineHorizontal(line []*block.Block) bool {
	currentPos := len(line) - 1
	currentPosStack := 0
	currentPosValue := 0
	moved := false

	insertBlock := func(b *block.Block) {
		if b == nil {
			return
		}

		if currentPosStack == 0 {
			currentPosValue += b.Number
			currentPosStack += 1
			if b.MoveHorizontal(currentPos) {
				moved = true
			}
			return
		}

		if currentPosStack == 1 && currentPosValue == b.GetValue() {

			currentPosValue += b.Number
			currentPosStack += 1
			if b.MoveHorizontal(currentPos) {
				moved = true
			}
			return

		}

		currentPos -= 1
		currentPosValue = b.Number
		currentPosStack = 1
		if b.MoveHorizontal(currentPos) {
			moved = true
		}
		return

	}

	for i := len(line) - 1; i >= 0; i-- {
		insertBlock(line[i])
	}
	return moved
}

func (p *Board) moveLineHorizontalReverse(line []*block.Block) bool {
	currentPos := 0
	currentPosStack := 0
	currentPosValue := 0
	moved := false

	insertBlock := func(b *block.Block) {
		if b == nil {
			return
		}

		if currentPosStack == 0 {
			currentPosValue += b.Number
			currentPosStack += 1
			if b.MoveHorizontal(currentPos) {
				moved = true
			}
			return
		}

		if currentPosStack == 1 && currentPosValue == b.GetValue() {

			currentPosValue += b.Number
			currentPosStack += 1
			if b.MoveHorizontal(currentPos) {
				moved = true
			}
			return

		}

		currentPos += 1
		currentPosValue = b.Number
		currentPosStack = 1
		if b.MoveHorizontal(currentPos) {
			moved = true
		}
		return

	}

	for i := 0; i < len(line); i++ {
		insertBlock(line[i])
	}

	return moved
}
