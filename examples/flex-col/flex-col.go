package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/shashanktomar/sprinkles/examples"
	"github.com/shashanktomar/sprinkles/flex"
)

type Bubble struct {
	layout *flex.Container
}

func New() Bubble {
	boxOne := examples.NewBox()
	boxTwo := examples.NewBox()
	boxThree := examples.NewBox()

	layout := flex.NewContainer(flex.Column).
		AddBox(boxOne, flex.NewStyle().Flex(1)).
		AddBox(boxTwo, flex.NewStyle().Flex(3)).
		AddBox(boxThree, flex.NewStyle().Flex(1))

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
