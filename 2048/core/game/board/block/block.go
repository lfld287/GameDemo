package block

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image/color"
	"strconv"
)

type Block struct {
	Number     int
	textColor  color.Color
	textFont   font.Face
	blockColor color.Color

	imageNumber int
	image       *ebiten.Image
	lPos        LogicalPosition

	pos        Position
	imageScale float64

	currentAction action
}

func NewBlock(number int, lPos LogicalPosition, blockColor color.Color, fontFace font.Face, textColor color.Color) *Block {
	res := &Block{
		Number:     number,
		blockColor: blockColor,
		textColor:  textColor,
		textFont:   fontFace,
		image:      nil,
		imageScale: 1,
		lPos:       lPos,
		pos: Position{
			X: float64(lPos.X),
			Y: float64(lPos.Y),
		},
		currentAction: nil,
	}
	res.initImage(number)
	res.Show(number)
	return res
}

func (p *Block) renderNumber() {
	str := strconv.Itoa(p.Number)

	var maxHeight fixed.Int26_6 = 0
	var totalWidth fixed.Int26_6 = 0

	for i := 0; i < len(str); i++ {
		r := str[i]
		bound, advanced, ok := p.textFont.GlyphBounds(rune(r))
		if !ok {
			panic("font not found for : " + string(r))
		}
		totalWidth += advanced

		height := bound.Max.Y - bound.Min.Y
		if maxHeight < height {
			maxHeight = height
		}
	}

	fmt.Printf("totalWidth: %v, maxHeight: %v\n", totalWidth.Round(), maxHeight.Round())

	//make sure the text is in the center
	var scale float64 = 1
	//width <= 75

	scale = 75 / float64(totalWidth.Round())

	// height <=75
	if float64(maxHeight.Round())*scale > 75 {
		scale = 75 / float64(maxHeight.Round())
	}

	finalWidth := float64(totalWidth.Round()) * scale
	finalHeight := float64(maxHeight.Round()) * scale

	var tx = 50 - finalWidth/2
	var ty = 50 + finalHeight/2

	fmt.Printf("scale: %v\n", scale)
	fmt.Printf("tx: %v, width: %v, ty: %v, height: %v\n", tx, finalWidth, ty, finalHeight)

	matrix := ebiten.GeoM{}
	matrix.Scale(scale, scale)
	matrix.Translate(tx, ty)
	text.DrawWithOptions(p.image, strconv.Itoa(p.Number), p.textFont, &ebiten.DrawImageOptions{
		GeoM: matrix,
	})
}

func (p *Block) initImage(number int) {
	if p.image == nil || p.imageNumber != number {
		p.imageNumber = number
		p.image = ebiten.NewImage(100, 100)
		p.image.Fill(p.blockColor)
		p.renderNumber()
	}
}

func (p *Block) Draw(screen *ebiten.Image, widthUnit, heightUnit float64) {
	p.initImage(p.Number)
	geoM := ebiten.GeoM{}
	geoM.Reset()
	geoM.Scale(widthUnit/100*p.imageScale, heightUnit/100*p.imageScale)
	geoM.Translate(p.pos.X*widthUnit, p.pos.Y*heightUnit)
	screen.DrawImage(p.image, &ebiten.DrawImageOptions{
		GeoM: geoM,
	})
}

func (p *Block) Update() (over bool) {
	if p.currentAction == nil {
		return true
	}
	over = p.currentAction.update(p)
	if over {
		p.currentAction = nil
	}
	return over
}

func (p *Block) GetPos() LogicalPosition {
	return p.lPos
}

func (p *Block) GetValue() int {
	return p.Number
}

func (p *Block) move(dst LogicalPosition) (moved bool) {
	if p.currentAction != nil {
		panic("current action is not nil")
	}
	if p.lPos == dst {
		return false
	}
	p.currentAction = newActionMove(p.lPos, dst)
	return true
}

func (p *Block) MoveVertical(row int) (moved bool) {
	return p.move(LogicalPosition{
		X: p.lPos.X,
		Y: row,
	})
}

func (p *Block) MoveHorizontal(col int) (moved bool) {
	return p.move(LogicalPosition{
		X: col,
		Y: p.lPos.Y,
	})
}

func (p *Block) Show(number int) {
	if p.currentAction != nil {
		panic("current action is not nil")
	}
	p.currentAction = newActionShow(number)
}
