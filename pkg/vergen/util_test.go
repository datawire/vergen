package vergen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_normalizeBranchName(t *testing.T) {
	tbl := []struct {
		name     string
		expected string
	}{
		{name: "foo/bar", expected: "foo-bar"},
		{name: "$$$foo//bar/", expected: "foo-bar"},
	}

	for _, tc := range tbl {
		assert.Equal(t, tc.expected, normalizeBranchName(tc.name))
	}
}
