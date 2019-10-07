package composite

import (
	"io/ioutil"
	"testing"

	"github.com/A1bemuth/deputy/types"
	"github.com/stretchr/testify/require"
)

func TestParsesCsproj(t *testing.T) {
	parser := New()
	const filename = "./test/example.csproj"
	content := getFileContent(filename)

	deps, err := parser.Parse(filename, content)
	require.Nil(t, err)
	require.Len(t, deps, 1)
	require.Equal(t,
		types.Dependency{Name: "Foo", Version: "1.0.0"},
		deps[0])
}

func getFileContent(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(f)
}
