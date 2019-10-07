package deputy

import (
	"github.com/A1bemuth/deputy/composite"
	"github.com/A1bemuth/deputy/types"
)

func NewParser() types.SelectiveParser {
	parser := composite.New()
	return &parser
}
