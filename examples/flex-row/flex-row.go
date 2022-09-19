package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shashanktomar/sprinkles/flex"
)

var boxStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())

// View you want to render in the box
type MyView struct {
	text   string
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
	text := lipgloss.Place(b.width-w, b.height-h, lipgloss.Center, lipgloss.Center, b.text)
	return b.style.Render(text)
}

type Bubble struct {
	layout *flex.Container
}

func New() Bubble {
	boxOne := &MyView{text: "box-one", style: boxStyle}
	boxTwo := &MyView{text: "box-two", style: boxStyle}
	boxThree := &MyView{text: "box-three", style: boxStyle}

	layout := flex.NewContainer(flex.Row).
		AddBox(boxOne, flex.NewBoxStyle(1)).
		AddBox(boxTwo, flex.NewBoxStyle(3)).
		AddBox(boxThree, flex.NewBoxStyle(2))

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
