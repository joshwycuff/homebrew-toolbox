package model

type FileItem struct {
	Name string
}

func (f FileItem) Title() string {
	return f.Name
}

func (f FileItem) Description() string {
	return ""
}

func (f FileItem) FilterValue() string {
	return f.Name
}
