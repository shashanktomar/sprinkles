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
	b.minSize = utils.Limit(0, b.maxSize, minSize)
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

// Adjust box size if it is outside min-max range
func (b *BoxStyle) limitSize() {
	b.basis = utils.Limit(b.minSize, b.maxSize, b.basis)
}

// Validate if changing the size break the bounds. If the new size is outside limits,
// limit it and return back actual adjusted size
func (b BoxStyle) limitSizeChange(size int) int {
	updatedSize := b.basis + size
	boundSize := utils.Limit(b.minSize, b.maxSize, updatedSize)
	if boundSize == updatedSize {
		return size
	} else {
		return boundSize - b.basis
	}
}

func (b BoxStyle) calculateFlexRatios(
	boxes []*BoxStyle,
) (growRatio float64, shrinkRatio float64) {
	totalGrow := 0.0
	totalShrink := 0.0
	for _, b := range boxes {
		totalGrow += b.grow
		totalShrink += b.shrink
	}

	if totalGrow == 0 {
		growRatio = 0
	} else {
		growRatio = b.grow / totalGrow
	}

	if totalShrink == 0 {
		shrinkRatio = 0
	} else {
		shrinkRatio = b.shrink / totalShrink
	}

	return growRatio, shrinkRatio
}
