package composite

import (
	"github.com/A1bemuth/deputy"
	"github.com/A1bemuth/deputy/csproj"
)

type CompositeParser struct {
	parsers []deputy.Parser
}

func New() CompositeParser {
	csprojParser := csproj.New()
	parsers := []deputy.Parser{&csprojParser}

	return CompositeParser{
		parsers: parsers,
	}
}

func (p *CompositeParser) Parse(filename string, content string) (deps []deputy.Dependency, err error) {
	parser := p.getParser(filename)
	if parser == nil {
		return nil, deputy.ErrExtensionNotSupported
	}

	return parser.Parse(content)
}

func (p *CompositeParser) getParser(filename string) deputy.Parser {
	for _, parser := range p.parsers {
		if parser.Accepts(filename) {
			return parser
		}
	}
	return nil
}
