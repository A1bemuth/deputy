package csproj

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsesValues(t *testing.T) {
	input := `<PackageReference Include="FluentAssertions" Version="5.4.1" />`
	parser := New()

	refs, err := parser.Parse(input)
	require.Nil(t, err)
	require.Equal(t, 1, len(refs))
	require.Equal(t, "FluentAssertions", refs[0].Name)
	require.Equal(t, "5.4.1", refs[0].Version)
}

func TestParsesCorrectCount(t *testing.T) {
	input := `
		<PackageReference Include="FluentAssertions" Version="5.4.1" />`
	testParsesCorrectCount(t, input, 1)
	input = `
		<PackageReference Include="foo" Version="1.0.0" />
		<PackageReference Include="bar" Version="2.0.0" />`
	testParsesCorrectCount(t, input, 2)
	input = `
		<ItemGroup>
			<PackageReference Include="FluentAssertions" Version="5.4.1" />
			<PackageReference Include="Microsoft.NET.Test.Sdk" Version="15.8.0" />
			<PackageReference Include="NSubstitute" Version="3.1.0" />
			<PackageReference Include="NUnit" Version="3.10.1" />
			<PackageReference Include="NUnit3TestAdapter" Version="3.10.0" />
		</ItemGroup>`
	testParsesCorrectCount(t, input, 5)
	input = `erence Include="FluentAssertions" Version="5.4.1" />
		<PackageReference Include="Microsoft.NET.Test.Sdk" Version="15.8.0" />`
	testParsesCorrectCount(t, input, 1)
	// input = `<PackageReference Include="Microsoft.NET.Test.Sdk">
	// 	<Version>15.8.0</Version>
	//  </PackageReference>`
	// testParsesCorrectCount(t, input, 1)
}

func testParsesCorrectCount(t *testing.T, input string, expected int) {
	parser := New()

	refs, err := parser.Parse(input)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, expected, len(refs))
}
