package flex

import (
	"math"

	"github.com/charmbracelet/lipgloss"
	"github.com/shashanktomar/sprinkles/utils"
)

// Direction is the direction in which flex behaviour is applied
type (
	Direction int8
)

const (
	Row Direction = iota
	Column
)

// Container holds the boxes and calculate their size when SetSize
// is called. The size is influenced by BoxStyle assigned to a Box
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

// NewContainer create a new container with a given Direction.
// Direction can row or column. This influence in what direction
// the flex behaviour will work.
func NewContainer(direction Direction) *Container {
	return &Container{
		direction: direction,
		boxes:     make([]*boxWithStyle, 0),
	}
}

// SetSize set the size of the container. This should be invoked
// on every size update, for example when you receive tea.WindowSizeMsg.
// Calling this method triggers the flex calculation on all child boxes.
func (c *Container) SetSize(width int, height int) {
	if c.direction == Row {
		c.mainSize = width
		c.crossSize = height
	} else {
		c.mainSize = height
		c.crossSize = width
	}
	c.calculate()
}

// AddBox adds a box to the container. Pass BoxStyle to set flex properties
// on the box.
func (c *Container) AddBox(box Box, style *BoxStyle) *Container {
	c.boxes = append(c.boxes, &boxWithStyle{
		box:   box,
		style: style,
	})
	return c
}

// View returns the rendered boxes. Use this in your View method to get the
// the styled string.
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

func (c *Container) calculate() {
	totalSizeToAdjust := c.mainSize
	for _, b := range c.boxes {
		b.style.limitSize()
		totalSizeToAdjust -= b.style.basis
	}

	totalAllocatedSize := 0
	for i, b := range c.boxes {
		remainingBoxes := boxStyles(c.boxes[i:])
		growRatio, shrinkRatio := b.style.calculateFlexRatios(remainingBoxes)
		ratio := 0.0
		if totalSizeToAdjust > 0 {
			ratio = growRatio
		} else {
			ratio = shrinkRatio
		}
		sizeToAdjust := int(math.Floor(ratio * float64(totalSizeToAdjust)))
		limitedSizeToAdjust := b.style.limitSizeChange(sizeToAdjust)
		totalSizeToAdjust -= limitedSizeToAdjust

		newSize := utils.Limit(b.style.minSize, b.style.maxSize, b.style.basis+limitedSizeToAdjust)
		totalAllocatedSize += newSize

		// If we missed few pixels because of fractions, allocate it to last box.
		if i == len(c.boxes)-1 && c.mainSize != totalAllocatedSize {
			newSize += c.mainSize - totalAllocatedSize
			newSize = utils.Limit(b.style.minSize, b.style.maxSize, newSize)
		}

		if c.direction == Row {
			b.box.SetSize(newSize, c.crossSize)
		} else {
			b.box.SetSize(c.crossSize, newSize)
		}
	}
}

func boxStyles(boxes []*boxWithStyle) []*BoxStyle {
	result := make([]*BoxStyle, len(boxes))
	for i, b := range boxes {
		result[i] = b.style
	}
	return result
}
