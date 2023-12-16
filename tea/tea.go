package tea

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/erikgeiser/promptkit/selection"
)

type model struct {
	items     []string
	selection *selection.Model[string]
}

func executeGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func listGitBranches() ([]string, error) {
	output, err := executeGitCommand("branch", "--all")
	if err != nil {
		return nil, err
	}
	branches := parseBranches(output)

	return branches, nil
}

func parseBranches(output string) []string {
	var branches []string
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		branch := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(branch, "remotes") {
			branches = append(branches, strings.Replace(branch, "remotes/", "", 1))
		} else if !strings.HasPrefix(branch, "*") {
			branches = append(branches, branch)
		}
	}
	return branches
}

func initialModel() *model {
	branches, err := listGitBranches()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return &model{
		items: branches,
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
			res, err := m.selection.Value()
			if err != nil {
				return m, tea.Quit
			}

			_, err = executeGitCommand("checkout", res)
			if err != nil {
				fmt.Println("Error:", err)
			}

			return m, tea.Quit

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

	s += "\nPress q to quit.\n"
	return s
}

func TryTea() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Could not start program:", err)
		os.Exit(1)
	}
}
