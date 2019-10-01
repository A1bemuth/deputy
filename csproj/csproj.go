package csproj

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"

	"github.com/A1bemuth/deputy"
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
	regex *regexp.Regexp
}

const packageRefRegex = `(<\s*PackageReference.+?((\/\s*>)|(<\s*\/\s*PackageReference\s*>)))`

// New creates an instance of Parser
func New() Parser {
	regex, err := regexp.Compile(packageRefRegex)
	if err != nil {
		panic(err)
	}
	return Parser{
		regex: regex,
	}
}

func (p *Parser) Accepts(filename string) bool {
	return strings.HasSuffix(filename, ".csproj")
}

func (p *Parser) Parse(content string) ([]deputy.Dependency, error) {
	matches := p.regex.FindAllString(content, -1)

	deps := make([]deputy.Dependency, 0)
	for _, match := range matches {
		ref, err := parsePackageRef(match)
		if err != nil {
			fmt.Printf("error parsing ref '%v', err: %v", match, err)
			continue
		}
		deps = append(deps, toDependency(*ref))
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

func toDependency(ref PackageReference) deputy.Dependency {
	return deputy.Dependency{
		Name:    ref.Name,
		Version: ref.Version,
	}
}
