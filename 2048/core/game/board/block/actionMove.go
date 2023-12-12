package block

var _ action = (*actionMove)(nil)

type actionMove struct {
	tick        int
	start       LogicalPosition
	destination LogicalPosition
}

func newActionMove(start LogicalPosition, destination LogicalPosition) *actionMove {
	return &actionMove{
		tick:        0,
		start:       start,
		destination: destination,
	}
}

func (p *actionMove) update(b *Block) (over bool) {
	p.tick += 1
	if p.tick >= moveTicks {
		b.pos = Position{
			X: float64(p.destination.X),
			Y: float64(p.destination.Y),
		}
		b.lPos = p.destination
		return true
	} else {
		progress := float64(p.tick) / float64(moveTicks)
		b.pos = Position{
			X: float64(p.start.X) + float64(p.destination.X-p.start.X)*progress,
			Y: float64(p.start.Y) + float64(p.destination.Y-p.start.Y)*progress,
		}
		return false
	}
}
