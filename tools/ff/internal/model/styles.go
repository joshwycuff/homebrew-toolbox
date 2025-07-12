package model

import "github.com/charmbracelet/lipgloss"

var (
	unselectedForegroundColor = lipgloss.Color("white")
	unselectedBackgroundColor = lipgloss.Color("black")

	selectedForegroundColor = lipgloss.Color("black")
	selectedBackgroundColor = lipgloss.Color("white")

	outerStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0).
		Margin(0)
	leftColumnStyle = lipgloss.NewStyle().
		Padding(0).
		Margin(0)
	rightColumnStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0).
		Margin(0)
	fileSearchStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0).
		Margin(0)
	fileFilterStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0).
		Margin(0)
	fileListStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0).
		Margin(0)
	fileItemStyle = lipgloss.NewStyle().
		Padding(0).
		Margin(0).
		Foreground(unselectedForegroundColor).
		Background(unselectedBackgroundColor)
	selectedItemStyle = lipgloss.NewStyle().
		Padding(0).
		Margin(0).
		Foreground(selectedForegroundColor).
		Background(selectedBackgroundColor).
		Bold(true)
)
