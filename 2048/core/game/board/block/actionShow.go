package block

var _ action = (*actionShow)(nil)

type actionShow struct {
	tick   int
	number int
}

func newActionShow(number int) *actionShow {
	return &actionShow{
		tick:   0,
		number: number,
	}
}

func (p *actionShow) update(b *Block) (over bool) {
	p.tick += 1
	b.Number = p.number
	if p.tick >= showTicks {
		b.imageScale = 1
		return true
	} else {
		//0 -> 1.15 ->1
		progress := float64(p.tick) / float64(showTicks)
		if progress < 0.5 {
			b.imageScale = 1 + (1.15-1)*progress*2
		} else {
			b.imageScale = 1 + (1.15-1)*(1-progress)*2
		}

		return false
	}
}
