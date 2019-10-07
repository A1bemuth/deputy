package types

import "errors"

var (
	ErrExtensionNotSupported = errors.New("File extension is not supported")
)

// Dependency represents known info about a single dependency
type Dependency struct {
	Name    string
	Version string
}

// Parser is a tool for extracting dependencies from certain fileTypes,
// given a file's content or it's part
type Parser interface {
	Accepts(filename string) bool
	Parse(content string) ([]Dependency, error)
}

type SelectiveParser interface {
	Parse(filename, content string) ([]Dependency, error)
}
