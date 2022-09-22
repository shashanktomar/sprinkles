package flex

import (
	"math"

	"github.com/charmbracelet/lipgloss"
	"github.com/shashanktomar/sprinkles/utils"
)

type (
	Direction int8
)

const (
	Row Direction = iota
	Column
)

type Container struct {
	direction Direction

	mainSize  int
	crossSize int

	boxes []*boxWithStyle
}

type boxWithStyle struct {
	box   Box
	style *BoxStyle
}

func NewContainer(direction Direction) *Container {
	return &Container{
		direction: direction,
		boxes:     make([]*boxWithStyle, 0),
	}
}

func (c *Container) SetSize(width int, height int) {
	if c.direction == Row {
		c.mainSize = width
		c.crossSize = height
	} else {
		c.mainSize = height
		c.crossSize = width
	}
	c.calculateBoxes()
}

func (c *Container) AddBox(box Box, style *BoxStyle) *Container {
	c.boxes = append(c.boxes, &boxWithStyle{
		box:   box,
		style: style,
	})
	return c
}

func (c *Container) View() string {
	text := make([]string, len(c.boxes))
	for i, b := range c.boxes {
		text[i] = b.box.View()
	}

	if c.direction == Row {
		return lipgloss.JoinHorizontal(lipgloss.Top, text...)
	} else {
		return lipgloss.JoinVertical(lipgloss.Left, text...)
	}
}

func (c *Container) calculateBoxes() {
	totalGrow := 0.0
	totalShrink := 0.0
	totalSizeToAdjust := c.mainSize
	for _, b := range c.boxes {
		totalGrow += b.style.grow
		totalShrink += b.style.shrink

		// adjust box size if it is outside min-max range
		b.style.basis = utils.Limit(b.style.minSize, b.style.maxSize, b.style.basis)

		totalSizeToAdjust -= b.style.basis
	}

	totalAllocatedSize := 0
	for i, b := range c.boxes {
		flex := 0.0
		if totalSizeToAdjust > 0 {
			flex = b.style.grow / totalGrow
		} else {
			flex = b.style.shrink / totalShrink
		}
		sizeToAdjust := int(math.Floor(flex * float64(totalSizeToAdjust)))
		updatedSize := b.style.basis + sizeToAdjust
		totalAllocatedSize += updatedSize

		// if we missed few pixels because of fractions, allocate it to last box
		if i == len(c.boxes)-1 && c.mainSize != totalAllocatedSize {
			updatedSize += c.mainSize - totalAllocatedSize
		}

		if c.direction == Row {
			b.box.SetSize(updatedSize, c.crossSize)
		} else {
			b.box.SetSize(c.crossSize, updatedSize)
		}
	}
}
