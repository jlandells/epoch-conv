package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	menu state = iota
	epochToHuman
	humanToEpoch
)

type model struct {
	cursor    int
	choices   []string
	state     state
	textInput textinput.Model
	output    string
	quitting  bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 20

	return model{
		choices:   []string{"Show current time in epoch (ms)", "Convert epoch to human-readable", "Convert human-readable to epoch", "Exit"},
		state:     menu,
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch m.state {
		case menu:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			case "enter":
				switch m.cursor {
				case 0:
					m.output = fmt.Sprintf("Current Epoch Time (ms): %d", time.Now().UnixNano()/int64(time.Millisecond))
				case 1:
					m.state = epochToHuman
					m.textInput.SetValue("")
					m.textInput.Focus()
					return m, textinput.Blink
				case 2:
					m.state = humanToEpoch
					m.textInput.SetValue("")
					m.textInput.Focus()
					return m, textinput.Blink
				case 3:
					m.quitting = true
					return m, tea.Quit
				}
			}
		case epochToHuman, humanToEpoch:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				input := m.textInput.Value()
				if m.state == epochToHuman {
					ms, err := strconv.ParseInt(input, 10, 64)
					if err != nil {
						m.output = "Invalid input"
					} else {
						m.output = fmt.Sprintf("Human-readable time: %s", time.Unix(0, ms*int64(time.Millisecond)).Format(time.RFC1123))
					}
				} else {
					layout := "2006-01-02 15:04:05"
					t, err := time.Parse(layout, input)
					if err != nil {
						m.output = "Invalid input"
					} else {
						m.output = fmt.Sprintf("Epoch Time (ms): %d", t.UnixNano()/int64(time.Millisecond))
					}
				}
				m.state = menu
				return m, nil
			}
			var cmd tea.Cmd
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What would you like to do?\n\n"

	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\nOutput:\n"
	s += m.output + "\n"

	if m.state == epochToHuman {
		s += fmt.Sprintf("\nEnter Epoch Time (ms):\n\n%s\n\n(press enter to confirm)", m.textInput.View())
	} else if m.state == humanToEpoch {
		s += fmt.Sprintf("\nEnter Human-readable time (YYYY-MM-DD HH:MM:SS):\n\n%s\n\n(press enter to confirm)", m.textInput.View())
	}

	if m.quitting {
		s += "\nBye!\n"
	}

	return s
}

func resetTerminal() {
	cmd := exec.Command("reset")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	// Ensure the terminal is reset on exit
	defer resetTerminal()

	// Create Bubble Tea program
	p := tea.NewProgram(initialModel(), tea.WithOutput(os.Stdout))

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
