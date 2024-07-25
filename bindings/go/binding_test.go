package tree_sitter_pyts_test

import (
	"testing"

	tree_sitter_pyts "github.com/ollipal/tree-sitter-pyts"
	tree_sitter "github.com/smacker/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_pyts.Language())
	if language == nil {
		t.Errorf("Error loading Pyts grammar")
	} else {
		t.Logf("Success loading Pyts grammar")
	}
}
