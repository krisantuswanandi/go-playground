package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erikgeiser/promptkit/selection"
)

type model struct {
	items     []string
	result    string
	selection *selection.Model[string]
}

func initialModel() *model {
	return &model{
		items: []string{"Carrots", "Celery", "Kohlrabi", "Cabbage", "Turnips", "Radishes", "Beets", "Potatoes", "Onions", "Garlic", "Shallots", "Leeks", "Scallions", "Chives", "Ginger", "Turmeric", "Rutabaga", "Parsnips", "Horseradish", "Wasabi", "Daikon", "Jicama", "Water chestnuts", "Yams", "Cassava", "Taro", "Lotus root", "Burdock", "Salsify", "Skirret", "Malanga", "Turnip root", "Crosne", "Ginger root"},
	}
}

func (m *model) Init() tea.Cmd {
	sel := selection.New("", m.items)
	sel.Filter = nil
	sel.PageSize = 7

	m.selection = selection.NewModel(sel)
	return m.selection.Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter", " ":
			if m.result != "" {
				return m, tea.Quit
			}

			res, err := m.selection.Value()
			if err != nil {
				return m, tea.Quit
			}

			m.result = res
			return m, nil

		default:
			_, cmd := m.selection.Update(msg)

			return m, cmd
		}

	}
	return m, nil
}

func (m *model) View() string {
	s := "What would you like to buy today?\n\n"

	s += m.selection.View()

	if m.result != "" {
		s += fmt.Sprintf("\nYou bought %s\n", m.result)
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Could not start program:", err)
		os.Exit(1)
	}
}
