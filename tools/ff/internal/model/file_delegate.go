package model

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"io"
)

type FileDelegate struct{}

func (d FileDelegate) Height() int                             { return 1 }
func (d FileDelegate) Spacing() int                            { return 0 }
func (d FileDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d FileDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	fileItem, ok := item.(FileItem)
	if !ok {
		return
	}

	s := fileItem.Name
	runes := []rune(s)
	l := len(runes)
	width := m.Width()
	if l > width {
		s = string(runes[:width-5]) + "..."
	}

	if index == m.Index() {
		s = fmt.Sprintf("> %s", s)
		_, _ = fmt.Fprint(w, selectedItemStyle.Render(s))
	} else {
		s = fmt.Sprintf("  %s", s)
		_, _ = fmt.Fprint(w, fileItemStyle.Render(s))
	}
}
