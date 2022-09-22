package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shashanktomar/sprinkles/flex"
)

var (
	boxStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	subtle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#8c8c8c"))
)

// View you want to render in the box
type MyView struct {
	background string
	color      lipgloss.Color
	width      int
	height     int
	style      lipgloss.Style
}

// Your view need to implement flex.Box interface
func (b *MyView) SetSize(width int, height int) {
	b.width = width
	b.height = height
	b.style = b.style.MaxWidth(b.width).MaxHeight(b.height)
}

// Your view need to implement flex.Box interface
func (b *MyView) View() string {
	w, h := b.style.GetFrameSize()
	width, height := b.width-w, b.height-h
	text := fmt.Sprintf("%d%s%d", width, subtle.Render("x"), height)
	result := lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.NewStyle().
			Padding(1, 1).
			Render(text),
		lipgloss.WithWhitespaceChars(b.background),
		lipgloss.WithWhitespaceForeground(b.color),
	)
	return b.style.Render(result)
}

type Bubble struct {
	layout *flex.Container
}

func New() Bubble {
	boxOne := &MyView{
		background: "│",
		color:      lipgloss.Color("#7dc4e4"),
		style:      boxStyle,
	}
	boxTwo := &MyView{
		background: "╰",
		color:      lipgloss.Color("#ed8796"),
		style:      boxStyle,
	}
	boxThree := &MyView{
		background: "╳",
		color:      lipgloss.Color("#c6a0f6"),
		width:      0,
		height:     0,
		style:      boxStyle,
	}
	boxFour := &MyView{
		background: "()",
		color:      lipgloss.Color("#b7bdf8"),
		style:      boxStyle,
	}
	boxFive := &MyView{
		background: "┼",
		color:      lipgloss.Color("#a6da95"),
		style:      boxStyle,
	}

	verticalLayoutOne := flex.NewContainer(flex.Column).
		AddBox(boxOne, flex.NewStyle().Flex(1)).
		AddBox(boxTwo, flex.NewStyle().Flex(5))
	verticalLayoutTwo := flex.NewContainer(flex.Column).
		AddBox(boxThree, flex.NewStyle().Flex(4)).
		AddBox(boxFour, flex.NewStyle().Flex(1))
	layout := flex.NewContainer(flex.Row).
		AddBox(verticalLayoutOne, flex.NewStyle().Flex(1)).
		AddBox(verticalLayoutTwo, flex.NewStyle().Flex(3)).
		AddBox(boxFive, flex.NewStyle().Flex(1.3))

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
