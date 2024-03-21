package tree_sitter_ejs_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-ejs"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_ejs.Language())
	if language == nil {
		t.Errorf("Error loading Ejs grammar")
	}
}
