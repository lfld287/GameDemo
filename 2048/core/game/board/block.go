package board

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
	number     int
	textColor  color.Color
	textFont   font.Face
	blockColor color.Color
	image      *ebiten.Image
}

func NewBlock(number int, blockColor color.Color, fontFace font.Face, textColor color.Color) *Block {
	res := &Block{
		number:     number,
		blockColor: blockColor,
		textColor:  textColor,
		textFont:   fontFace,
		image:      nil,
	}
	return res
}

func (p *Block) Render() *ebiten.Image {

	if p.image == nil {
		p.image = ebiten.NewImage(100, 100)
		p.image.Fill(p.blockColor)
		p.renderNumber()
	}

	return p.image
}

func (p *Block) renderNumber() {
	str := strconv.Itoa(p.number)

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
	text.DrawWithOptions(p.image, strconv.Itoa(p.number), normalFont, &ebiten.DrawImageOptions{
		GeoM: matrix,
	})
}
