package model

// Document is a Markdown document with YAML frontmatter.
type Document struct {
	Frontmatter map[string]any
	Body        string
}

// Item is one level-three spec item. Body excludes the heading line.
type Item struct {
	ID   string
	Tag  string
	Body string
}

// Zones holds module-relative path patterns for each editable zone.
type Zones struct {
	Protected  []string
	Evals      []string
	Interface  []string
	AppendOnly []string
	Preserve   []string
	Disposable []string
}
