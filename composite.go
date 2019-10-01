package deputy

import (
	"github.com/A1bemuth/deputy/csproj"
)

type CompositeParser struct {
	parsers []Parser
}

func New() CompositeParser {
	parsers := [1]Parser{csproj.New()}
}

func (p *CompositeParser) Parse(filename string, content string) (deps []Dependency, err error) {
	parser := p.getParser(filename)
	if parser == nil {
		return nil, ErrExtensionNotSupported
	}

	return parser.Parse(content)
}

func (p *CompositeParser) getParser(filename string) Parser {
	for _, parser := range p.parsers {
		if parser.Accepts(filename) {
			return parser
		}
	}
	return nil
}
