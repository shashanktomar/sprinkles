package examples

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	subtle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#8c8c8c"))
	boxStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
)

// View you want to render in the box
type Box struct {
	width  int
	height int
	style  lipgloss.Style
}

func NewBox() *Box {
	return &Box{style: boxStyle}
}

// Your view need to implement flex.Box interface
func (b *Box) SetSize(width int, height int) {
	b.width = width
	b.height = height
	b.style = b.style.MaxWidth(b.width).MaxHeight(b.height)
}

// Your view need to implement flex.Box interface
func (b *Box) View() string {
	if b.width == 0 || b.height == 0 {
		return ""
	}
	w, h := b.style.GetFrameSize()
	display := fmt.Sprintf("%d%s%d", b.width, subtle.Render("x"), b.height)
	text := lipgloss.Place(b.width-w, b.height-h, lipgloss.Center, lipgloss.Center, display)
	return b.style.Render(text)
}
