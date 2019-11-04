package csproj

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptsCsproj(t *testing.T) {
	parser := Parser{}
	testAccepts := func(name string, expected bool) func(t *testing.T) {
		return func(t *testing.T) {
			accepts := parser.Accepts(name)
			assert.Equal(t, expected, accepts)
		}
	}

	t.Run("Csproj1", testAccepts("Foo.csproj", true))
	t.Run("Csproj2", testAccepts("Foo bar.csproj", true))
	t.Run("Sln1", testAccepts("Foo.sln", false))
	t.Run("Sln2", testAccepts("Foo.csproj.sln", false))
}

func TestParsesValues(t *testing.T) {
	input := `<PackageReference Include="FluentAssertions" Version="5.4.1" />`
	parser := Parser{}

	refs, err := parser.Parse(input)
	require.Nil(t, err)
	require.Equal(t, 1, len(refs))
	require.Equal(t, "FluentAssertions", refs[0].Name)
	require.Equal(t, "5.4.1", refs[0].Version)
}

func TestParsesCorrectCount(t *testing.T) {
	input := `
		<PackageReference Include="FluentAssertions" Version="5.4.1" />`
	t.Run("Single entry", testParsesCorrectCount(input, 1))

	input = `
		<PackageReference 
			Include="FluentAssertions" 
			Version="5.4.1" />`
	t.Run("Single entry, multiline", testParsesCorrectCount(input, 1))

	input = `
		<PackageReference Include="foo" Version="1.0.0" />
		<PackageReference Include="bar" Version="2.0.0" />`
	t.Run("Two entries", testParsesCorrectCount(input, 2))

	input = `
		<ItemGroup>
			<PackageReference Include="FluentAssertions" Version="5.4.1" />
			<PackageReference Include="Microsoft.NET.Test.Sdk" Version="15.8.0" />
			<PackageReference Include="NSubstitute" Version="3.1.0" />
			<PackageReference Include="NUnit" Version="3.10.1" />
			<PackageReference Include="NUnit3TestAdapter" Version="3.10.0" />
		</ItemGroup>`
	t.Run("Whole ItemGroup", testParsesCorrectCount(input, 5))

	input = `erence Include="FluentAssertions" Version="5.4.1" />
		<PackageReference Include="Microsoft.NET.Test.Sdk" Version="15.8.0" />`
	t.Run("Partial XML", testParsesCorrectCount(input, 1))

	input = `
		<ItemGroup>
    		<PackageReference Update="NETStandard.Library" Version="2.0.1" />
		</ItemGroup>`
	t.Run("No package name", testParsesCorrectCount(input, 0))
	//TODO
	// input = `<PackageReference Include="Microsoft.NET.Test.Sdk">
	// 	<Version>15.8.0</Version>
	//  </PackageReference>`
	// testParsesCorrectCount(t, input, 1)
}

func testParsesCorrectCount(input string, expected int) func(t *testing.T) {
	return func(t *testing.T) {
		parser := Parser{}

		refs, err := parser.Parse(input)
		if !assert.Nil(t, err) {
			return
		}
		assert.Equal(t, expected, len(refs))
	}
}
