package flex

import (
	"math"

	"github.com/shashanktomar/sprinkles/utils"
)

const None = 0

type Box interface {
	SetSize(int, int)
	View() string
}

type BoxStyle struct {
	grow    float64
	shrink  float64
	basis   int
	minSize int
	maxSize int
}

func NewStyle() *BoxStyle {
	b := &BoxStyle{
		maxSize: math.MaxInt,
	}
	b.FlexDefault()
	return b
}

func (b *BoxStyle) Grow(grow float64) *BoxStyle {
	b.grow = math.Max(0, grow)
	return b
}

func (b *BoxStyle) Shrink(shrink float64) *BoxStyle {
	b.shrink = math.Max(0, shrink)
	return b
}

func (b *BoxStyle) Basis(basis int) *BoxStyle {
	b.basis = utils.Max(0, basis)
	return b
}

func (b *BoxStyle) MinSize(minSize int) *BoxStyle {
	b.minSize = utils.Limit(0, b.maxSize, b.minSize)
	return b
}

func (b *BoxStyle) MaxSize(maxSize int) *BoxStyle {
	b.maxSize = utils.Limit(b.minSize, math.MaxInt, maxSize)
	return b
}

func (b *BoxStyle) FlexDefault() *BoxStyle {
	return b.FlexCustom(1, 1, None)
}

func (b *BoxStyle) Flex(grow float64) *BoxStyle {
	return b.FlexCustom(grow, 1, None)
}

func (b *BoxStyle) FlexAuto(basis int) *BoxStyle {
	return b.FlexCustom(1, 1, basis)
}

func (b *BoxStyle) FlexCustom(grow float64, shrink float64, basis int) *BoxStyle {
	b.Grow(grow)
	b.Shrink(shrink)
	b.Basis(basis)
	return b
}

func (b *BoxStyle) FlexNone(basis int) *BoxStyle {
	return b.FlexCustom(None, None, basis)
}
