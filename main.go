package main

import (
	// "fmt"
	"log"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct{
  active int
  views []tea.Model
}

func main(){
  p := tea.NewProgram(NewModel())
  _, err := p.Run()
  if err != nil {
    log.Fatal(err)
  } 
}

func NewModel() *Model {
  return &Model{
    views: []tea.Model{NewActive()},
  }
}

func (m Model) Init() tea.Cmd {
  //Iterate through and call Init() on all the views. Return all the cmds
  var cmds []tea.Cmd
  for _, view := range m.views{
    cmds = append(cmds, view.Init())
  }
  return tea.Batch(cmds...)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd){

  var cmd tea.Cmd
  m.views[m.active], cmd = m.views[m.active].Update(msg)

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    }
  }
  return m, cmd
}

func (m Model) View() string {
  return m.views[m.active].View()
}
