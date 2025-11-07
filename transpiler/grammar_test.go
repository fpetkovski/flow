package transpiler

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrammar(t *testing.T) {
	casesContent, err := os.ReadFile("testdata/valid")
	require.NoError(t, err)

	cases := []string{""}
	scanner := bufio.NewScanner(bytes.NewReader(casesContent))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}
		if len(line) == 0 {
			cases = append(cases, "")
			continue
		}
		n := len(cases) - 1
		cases[n] = cases[n] + " " + line
	}

	for i := range cases {
		t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
			_, errs := parseQuery(cases[i])
			require.Empty(t, len(errs))
		})
	}
}
