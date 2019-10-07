package composite

import (
	"github.com/A1bemuth/deputy/csproj"
	. "github.com/A1bemuth/deputy/types"
)

type CompositeParser struct {
	parsers []Parser
}

func New() CompositeParser {
	csprojParser := csproj.New()
	parsers := []Parser{&csprojParser}

	return CompositeParser{
		parsers: parsers,
	}
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
