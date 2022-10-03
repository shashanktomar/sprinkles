package flex

import (
	"math"

	"github.com/shashanktomar/sprinkles/utils"
)

const None = 0

// Box is placed in a Container along with other boxes.
// On calculating flex width or height, SetSize is invoked
// on the Box.
// View is called just before rendering to get the string
// that need to be rendered.
type Box interface {
	SetSize(int, int)
	View() string
}

// BoxStyle is used to set properties which influence
// the flex behaviour of a box
type BoxStyle struct {
	grow    float64
	shrink  float64
	basis   int
	minSize int
	maxSize int
}

// NewStyle returns a BoxStyle with default values.
func NewStyle() *BoxStyle {
	b := &BoxStyle{
		maxSize: math.MaxInt,
	}
	b.FlexDefault()
	return b
}

// MinSize set the minimum allowed size for the box. A view
// can not shrink beyond the minSize
// If value is provided outside 0 - maxSize range, it is
// shifted to either 0 or maxSize.
func (b *BoxStyle) MinSize(minSize int) *BoxStyle {
	b.minSize = utils.Limit(0, b.maxSize, minSize)
	return b
}

// MaxSize set the maximum allowed size for the box. A view
// can not grow beyond maxSize.
// If value is less than minSize, it is shifted to minSize
func (b *BoxStyle) MaxSize(maxSize int) *BoxStyle {
	b.maxSize = utils.Limit(b.minSize, math.MaxInt, maxSize)
	return b
}

// FlexDefault set the default value for grow, shrink and basis.
// The default values are 1, 1, 0. This means that the box will
// have no starting basis and will start from 0 size. It will then
// grow with a proportion of 1. Shrink will have no effect as there
// is no basis to start with.
func (b *BoxStyle) FlexDefault() *BoxStyle {
	return b.FlexCustom(1, 1, None)
}

// Flex is expected to be the most commonly used method and only
// allows setting a grow value. It defaults shrink and basis to
// 1, 0.
func (b *BoxStyle) Flex(grow float64) *BoxStyle {
	return b.FlexCustom(grow, 1, None)
}

// FlexAuto can be used to set only the basis and defaults grow and shrink
// to 1, 1.
func (b *BoxStyle) FlexAuto(basis int) *BoxStyle {
	return b.FlexCustom(1, 1, basis)
}

// FlexCustom can be used to set all three values, grow, shrink and basis.
func (b *BoxStyle) FlexCustom(grow float64, shrink float64, basis int) *BoxStyle {
	b.setGrow(grow)
	b.setShrink(shrink)
	b.setBasis(basis)
	return b
}

// FlexNone can be used to set only the basis and defaults grow and shrink
// to 0, 0. This is a way to disable the flex behaviour on a box.
func (b *BoxStyle) FlexNone(basis int) *BoxStyle {
	return b.FlexCustom(None, None, basis)
}

func (b *BoxStyle) setGrow(grow float64) *BoxStyle {
	b.grow = math.Max(0, grow)
	return b
}

func (b *BoxStyle) setShrink(shrink float64) *BoxStyle {
	b.shrink = math.Max(0, shrink)
	return b
}

func (b *BoxStyle) setBasis(basis int) *BoxStyle {
	b.basis = utils.Max(0, basis)
	return b
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
