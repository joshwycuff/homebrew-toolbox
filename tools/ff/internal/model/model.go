package model

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joshwycuff/homebrew-toolbox/tools/ff/internal/preview"
	"github.com/joshwycuff/homebrew-toolbox/tools/ff/internal/search"
	"github.com/sahilm/fuzzy"
	"os"
)

const (
	focusFileSearch  = 1
	focusFileFilter  = 2
	focusFileList    = 3
	focusFilePreview = 4

	minOuterWidth  = 32
	minOuterHeight = 32

	borderSize = 2 // 1 top + 1 bottom or 1 left + 1 right
)

type Model struct {
	width  int
	height int
	focus  int

	debug string

	fileSearch    textinput.Model
	prevSearch    string
	err           string
	searchResults []string

	fileFilter     textinput.Model
	showFileFilter bool
	prevFilter     string

	fileList list.Model

	fileView viewport.Model

	// Return values
	Selected []string
	ExitCode int
}

func New(initialQuery string) Model {

	fileSearch := textinput.New()
	fileSearch.CharLimit = 128
	fileSearch.Placeholder = "File Search"
	fileSearch.Prompt = ""
	fileSearch.SetValue(initialQuery)
	fileSearch.Focus()

	fileFilter := textinput.New()
	fileFilter.CharLimit = 128
	fileFilter.Placeholder = "File Filter"
	fileFilter.Prompt = ""
	fileFilter.SetValue("")

	fileList := list.New(nil, FileDelegate{}, 0, 0)
	fileList.SetShowFilter(false)
	fileList.SetShowTitle(false)
	fileList.SetShowStatusBar(false)
	fileList.SetShowHelp(false)

	fileView := viewport.New(0, 0)

	return Model{
		focus:      focusFileSearch,
		fileSearch: fileSearch,
		fileFilter: fileFilter,
		fileList:   fileList,
		fileView:   fileView,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			m.focus = rotateFocus(m)
		case tea.KeyCtrlN:
			m.fileList.Select(min(m.fileList.Index()+1, len(m.fileList.Items())-1))
		case tea.KeyCtrlP:
			m.fileList.Select(max(m.fileList.Index()-1, 0))
		case tea.KeyCtrlF:
			m.showFileFilter = !m.showFileFilter
		case tea.KeyCtrlD:
			m.debug = ""
		case tea.KeyCtrlR:
			if m.focus == focusFileSearch {
				m.prevSearch = ""
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			m.ExitCode = 1
			return m, tea.Quit
		case tea.KeyCtrlY:
			m.Selected = []string{m.fileList.SelectedItem().(FileItem).Name}
			m.ExitCode = 0
			return m, tea.Quit
		case tea.KeyEnter:
			m.Selected = make([]string, len(m.fileList.Items()))
			for i, item := range m.fileList.Items() {
				if fileItem, ok := item.(FileItem); ok {
					m.Selected[i] = fileItem.Name
				} else {
					m.Selected[i] = ""
				}
			}
			m.ExitCode = 0
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	updateSearch := m.fileSearch.Value() != m.prevSearch
	updateFilter := updateSearch || (m.fileFilter.Value() != m.prevFilter)

	if updateSearch {
		m.fileList.Select(0)
		m.prevSearch = m.fileSearch.Value()

		files, stderr, err := search.SearchFiles(m.fileSearch.Value())
		if err != nil {
			m.err = err.Error()
		}
		if stderr != "" {
			m.err = stderr
		}
		m.err = ""
		m.searchResults = files

		items := make([]list.Item, len(files))
		for i, file := range files {
			items[i] = FileItem{Name: file}
		}
		m.fileList.SetItems(items)
	}

	if updateFilter {
		m.fileList.Select(0)
		m.prevFilter = m.fileFilter.Value()

		if m.fileFilter.Value() == "" {
			items := make([]list.Item, len(m.searchResults))
			for i, filename := range m.searchResults {
				items[i] = FileItem{Name: filename}
			}
			m.fileList.SetItems(items)
		} else {
			matches := fuzzy.Find(m.fileFilter.Value(), m.searchResults)
			filteredItems := make([]list.Item, len(matches))
			for i, match := range matches {
				filename := match.Str
				filteredItems[i] = FileItem{Name: filename}
			}
			m.fileList.SetItems(filteredItems)
		}
	}

	if m.focus == focusFileSearch {
		m.fileSearch.Focus()
		var fileSearchCmd tea.Cmd
		m.fileSearch, fileSearchCmd = m.fileSearch.Update(msg)
		cmds = append(cmds, fileSearchCmd)
	} else if m.focus == focusFileFilter {
		m.fileFilter.Focus()
		var fileFilterCmd tea.Cmd
		m.fileFilter, fileFilterCmd = m.fileFilter.Update(msg)
		cmds = append(cmds, fileFilterCmd)
	} else if m.focus == focusFileList {
		var fileListCmd tea.Cmd
		m.fileList, fileListCmd = m.fileList.Update(msg)
		cmds = append(cmds, fileListCmd)
	} else if m.focus == focusFilePreview {
		var fileViewCmd tea.Cmd
		m.fileView, fileViewCmd = m.fileView.Update(msg)
		cmds = append(cmds, fileViewCmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	outerWidth := max(minOuterWidth, m.width-borderSize)
	outerHeight := max(minOuterHeight, m.height-borderSize)

	outerStyle := outerStyle.
		Width(outerWidth).
		Height(outerHeight)

	leftColumnWidth := outerWidth/2 - borderSize
	leftColumnHeight := outerHeight - borderSize

	rightColumnWidth := leftColumnWidth
	rightColumnHeight := leftColumnHeight

	leftColumn := renderLeftColumn(m, leftColumnWidth, leftColumnHeight)

	rightColumn := renderRightColumn(m, rightColumnWidth, rightColumnHeight)

	columns := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, rightColumn)

	return outerStyle.Render(columns)
}

func renderLeftColumn(m Model, leftColumnWidth, leftColumnHeight int) string {
	leftColumnStyle_ := leftColumnStyle.
		Width(leftColumnWidth).
		Height(leftColumnHeight)

	fileSearchStyleWidth := leftColumnWidth - borderSize
	fileSearchWidth := fileSearchStyleWidth - borderSize
	fileSearchHeight := numLines(m.fileSearch.Value(), fileSearchWidth)
	fileSearchStyleHeight := fileSearchHeight

	fileSearchStyle := fileSearchStyle.
		Width(fileSearchStyleWidth).
		Height(fileSearchStyleHeight)
	m.fileSearch.Width = fileSearchWidth

	if m.focus == focusFileSearch {
		m.fileSearch.Focus()
		fileSearchStyle = fileSearchStyle.
			BorderForeground(selectedForegroundColor).
			BorderBackground(selectedBackgroundColor)
	} else {
		m.fileSearch.Blur()
		fileSearchStyle = fileSearchStyle.
			BorderForeground(unselectedForegroundColor).
			BorderBackground(unselectedBackgroundColor)
	}

	fileFilterStyleWidth := leftColumnWidth - borderSize
	fileFilterWidth := fileFilterStyleWidth - borderSize
	fileFilterHeight := numLines(m.fileFilter.Value(), fileFilterWidth)
	fileFilterStyleHeight := fileFilterHeight

	fileFilterStyle := fileFilterStyle.
		Width(fileFilterStyleWidth).
		Height(fileFilterStyleHeight)
	m.fileFilter.Width = fileFilterWidth

	if m.focus == focusFileFilter {
		m.fileFilter.Focus()
	} else {
		m.fileFilter.Blur()
	}

	fileListStyle := fileListStyle.
		Width(leftColumnWidth - borderSize)

	if m.showFileFilter {
		fileListStyle = fileListStyle.
			Height(leftColumnHeight - fileSearchStyle.GetHeight() - fileFilterStyle.GetHeight() - 2*borderSize)
	} else {
		fileListStyle = fileListStyle.
			Height(leftColumnHeight - fileSearchStyle.GetHeight() - borderSize)
	}

	m.fileList.SetWidth(fileListStyle.GetWidth() - borderSize)
	m.fileList.SetHeight(fileListStyle.GetHeight() - borderSize)

	var innerLeftColumn string
	if m.showFileFilter {
		innerLeftColumn = lipgloss.JoinVertical(
			lipgloss.Left,
			fileSearchStyle.Render(m.fileSearch.View()),
			fileFilterStyle.Render(m.fileFilter.View()),
			fileListStyle.Render(m.fileList.View()),
		)
	} else {
		innerLeftColumn = lipgloss.JoinVertical(
			lipgloss.Left,
			fileSearchStyle.Render(m.fileSearch.View()),
			fileListStyle.Render(m.fileList.View()),
		)
	}

	return leftColumnStyle_.Render(innerLeftColumn)
}

func renderRightColumn(m Model, rightColumnWidth, rightColumnHeight int) string {
	rightColumnStyle := rightColumnStyle.
		Width(rightColumnWidth).
		Height(rightColumnHeight)

	if m.debug != "" {
		return rightColumnStyle.Render(m.debug)
	}

	if m.err != "" {
		return rightColumnStyle.Render("Error: " + m.err)
	}

	selected, ok := m.fileList.SelectedItem().(FileItem)
	if !ok {
		return rightColumnStyle.Render("No file selected")
	}

	filename := selected.Name
	if !fileExists(filename) {
		return rightColumnStyle.Render("File does not exist: " + filename)
	}

	content, err := preview.Preview(filename)
	if err != nil {
		return rightColumnStyle.Render("Error previewing file: " + err.Error())
	}

	m.fileView.Width = rightColumnStyle.GetWidth() - borderSize
	m.fileView.Height = rightColumnStyle.GetHeight() - borderSize
	m.fileView.SetContent(content)

	return rightColumnStyle.Render(m.fileView.View())
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func numLines(text string, width int) int {
	if width <= 0 {
		return 0
	}
	runes := []rune(text) // handle Unicode safely
	length := len(runes)
	if length == 0 {
		return 1
	}
	lines := length / width
	if length%width != 0 {
		lines++
	}
	return lines
}

func rotateFocus(m Model) int {
	switch m.focus {
	case focusFileSearch:
		if m.showFileFilter {
			return focusFileFilter
		} else {
			return focusFileList
		}
	case focusFileFilter:
		return focusFileList
	case focusFileList:
		return focusFilePreview
	case focusFilePreview:
		return focusFileSearch
	}
	return focusFileSearch
}
