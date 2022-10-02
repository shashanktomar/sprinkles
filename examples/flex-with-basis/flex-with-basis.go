package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shashanktomar/sprinkles/flex"
)

var (
	boxStyle  = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	subtle    = lipgloss.Color("#383838")
	highlight = lipgloss.Color("#7D56F4")
)

// View you want to render in the box
type MyView struct {
	basis  int
	width  int
	height int
	style  lipgloss.Style
}

// Your view need to implement flex.Box interface
func (b *MyView) SetSize(width int, height int) {
	b.width = width
	b.height = height
}

// Your view need to implement flex.Box interface
func (b *MyView) View() string {
	w, h := b.style.GetFrameSize()
	return b.style.Render(b.generateText(b.width-w, b.height-h))
}

func (b *MyView) generateText(width int, height int) string {
	extraSizeOnEachSide := (width - b.basis) / 2
	extendedText := ""
	if extraSizeOnEachSide > 0 {
		extendedText = lipgloss.NewStyle().
			Foreground(subtle).
			Render(strings.Repeat("~", extraSizeOnEachSide))
	}

	var sb strings.Builder
	sb.WriteString(extendedText)
	sb.WriteString(strings.Repeat("#", (b.basis-2)/2))
	sizeText := lipgloss.NewStyle().
		Foreground(highlight).
		Render(fmt.Sprint(b.basis))
	sb.WriteString(sizeText)
	sb.WriteString(strings.Repeat("#", (b.basis-2)/2))
	sb.WriteString(extendedText)

	text := lipgloss.NewStyle().MaxWidth(width).Render(sb.String())

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, text)
}

type Bubble struct {
	layout *flex.Container
}

func New() Bubble {
	boxOne := &MyView{basis: 20, style: boxStyle}
	boxTwo := &MyView{basis: 30, style: boxStyle}
	boxThree := &MyView{basis: 20, style: boxStyle}

	layout := flex.NewContainer(flex.Row).
		AddBox(boxOne, flex.NewStyle().FlexCustom(2, 1, 20)).
		AddBox(boxTwo, flex.NewStyle().FlexCustom(1, 1, 30)).
		AddBox(boxThree, flex.NewStyle().FlexCustom(1, 1, 20))

	return Bubble{
		layout: layout,
	}
}

func (b Bubble) Init() tea.Cmd {
	return nil
}

func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.layout.SetSize(msg.Width, msg.Height)
		return b, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return b, tea.Quit
		}
	}
	return b, nil
}

func (b Bubble) View() string {
	return b.layout.View()
}

func main() {
	b := New()
	p := tea.NewProgram(b, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
