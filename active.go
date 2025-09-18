package main

import (
	// "github.com/charmbracelet/bubbles/list"
	"fmt"
	"time"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  "github.com/charmbracelet/huh"
)

type active struct{
  timer timer.Model
  progress progress.Model
  form *huh.Form
  // list  list.Model
}

func NewActive() active {
  return active{
    timer: timer.New(time.Duration(15 * time.Second)),
    progress: progress.New(
      progress.WithDefaultGradient(),
      progress.WithWidth(100),
      // progress.WithSpringOptions(20, 20),
    ),
  }
}

func (a active) Init() tea.Cmd {
  return tea.Batch(a.timer.Init(), a.progress.Init())
}

func (a active) Update(msg tea.Msg) (tea.Model, tea.Cmd){
  var cmd tea.Cmd
  var m tea.Model
  var cmds []tea.Cmd

  // switch msg.(type) {
  // case timer.TickMsg: 
  //   cmds = append(cmds, a.progress.SetPercent(float64(15-a.timer.Timeout.Seconds())/15))
  // }
  m, cmd = a.progress.Update(msg)
  a.progress = m.(progress.Model)
  cmds = append(cmds, cmd)
  
  a.timer, cmd = a.timer.Update(msg)
  cmds = append(cmds, cmd)
  return a, tea.Batch(cmds...)
}

func (a active) View() string {
  return lipgloss.JoinVertical(lipgloss.Left, 
    "Your timer is running!", 
    fmt.Sprintf("Next break starts in %s", a.timer.View()),
    a.progress.ViewAs(float64(15-a.timer.Timeout.Seconds())/15),
  ) 
}
