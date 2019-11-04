package csproj

import (
	"encoding/xml"
	"regexp"
	"strings"

	"github.com/A1bemuth/deputy/types"
)

// PackageReference represents a ref to a nuget package
type PackageReference struct {
	Name    string `xml:"Include,attr"`
	Version string `xml:"Version,attr"`
}

// ItemGroup represents a collection of msbuild items
type ItemGroup struct {
	XMLName xml.Name           `xml:"ItemGroup"`
	Refs    []PackageReference `xml:"PackageReference"`
}

// Parser helps extract package references from a csproj file or its part
type Parser struct {
}

const ECOSYSTEM = "nuget"
const PACKAGE_REF_REGEXP = `(?s)(<\s*PackageReference.+?((\/\s*>)|(<\s*\/\s*PackageReference\s*>)))`

var referenceRegex = regexp.MustCompile(PACKAGE_REF_REGEXP)

func (p *Parser) Accepts(filename string) bool {
	return strings.HasSuffix(filename, ".csproj")
}

func (p *Parser) Parse(content string) ([]types.Dependency, error) {
	matches := referenceRegex.FindAllString(content, -1)

	deps := make([]types.Dependency, 0)
	for _, match := range matches {
		ref, err := parsePackageRef(match)
		if err != nil {
			return nil, err
		}
		if ref.Name != "" && ref.Version != "" {
			deps = append(deps, toDependency(*ref))
		}
	}

	return deps, nil
}

func parsePackageRef(xmlStr string) (*PackageReference, error) {
	ref := PackageReference{}
	if err := xml.Unmarshal([]byte(xmlStr), &ref); err != nil {
		return nil, err
	}
	return &ref, nil
}

func toDependency(ref PackageReference) types.Dependency {
	return types.Dependency{
		Ecosystem: ECOSYSTEM,
		Name:      ref.Name,
		Version:   ref.Version,
	}
}
