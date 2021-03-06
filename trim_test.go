package sweets

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Trim(t *testing.T) {
	t.Parallel()

	cab := os.DirFS("testdata")

	table := []struct {
		in  string
		exp string
	}{
		{in: "a.txt", exp: "line a\nline b\nline c"},
		{in: "b.txt", exp: "  line a\nline b\n  line c"},
		{in: "c.txt", exp: "line a\nline b\nline c"},
		{in: "d.txt", exp: "line a\n  line b\n\nline c"},
		{in: "e.txt", exp: "package main\n\nimport (\n\t\"io\"\n\t\"os\"\n)\n\nfunc main() {\n\twriteSomething(os.Stdout)\n\t// Hello\n}\n\nfunc writeSomething(w io.Writer) {\n\tw.Write([]byte(\"Hello\"))\n}\n"},
		// {in: "f.txt", exp: "line a\n  line b\n\nline c"},
	}

	for _, tt := range table {
		t.Run(tt.in, func(t *testing.T) {
			r := require.New(t)

			b, err := fs.ReadFile(cab, tt.in)
			r.NoError(err)

			act := TrimLeftSpace(string(b))
			r.Equal(tt.exp, act)

		})
	}
}
