package deputy

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
	Parse(content string) (deps []Dependency, err error)
}
