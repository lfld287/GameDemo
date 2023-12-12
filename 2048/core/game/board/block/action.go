package block

const (
	moveTicks = 12
	showTicks = 6
)

type action interface {
	update(b *Block) (over bool)
}
