package flex

import (
	"math"

	"github.com/charmbracelet/lipgloss"
)

type Direction int8

const (
	Row Direction = iota
	Column
)

type Container struct {
	direction Direction

	width  int
	height int

	boxes []*boxWrapper
}

// TODO: Rename this
type boxWrapper struct {
	box    Box
	config boxConfig
}

func NewContainer(direction Direction) *Container {
	return &Container{
		direction: direction,
	}
}

func (c *Container) SetSize(width int, height int) {
	c.width = width
	if c.width < 0 {
		c.width = 0
	}

	c.height = height
	if c.height < 0 {
		c.height = 0
	}

	c.calculateBoxes()
}

func (c *Container) AddBox(box Box, config boxConfig) *Container {
	// TODO: validate config
	c.boxes = append(c.boxes, &boxWrapper{
		box:    box,
		config: config,
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
	totalRatio := 0
	for _, b := range c.boxes {
		totalRatio += b.config.ratio
	}

	for _, b := range c.boxes {
		sizeFraction := float64(b.config.ratio) / float64(totalRatio)
		if c.direction == Row {
			width := int(math.Floor(sizeFraction * float64(c.width)))
			b.box.SetSize(width, c.height)
		} else {
			height := int(math.Floor(sizeFraction * float64(c.height)))
			b.box.SetSize(c.width, height)
		}
	}
}
