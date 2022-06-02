package jsdoc_test

import (
	"context"
	"testing"

	jsdoc "github.com/DaivikDave/tree-sitter-jsdoc/bindings/go"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/tj/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	code := `/**
 	* Blend two colors together.
 	* @param {string} color1 - The first color, in hexadecimal format.
 	* @param {string} color2 - The second color, in hexadecimal format.
 	* @return {string} The blended color.
 	*/
	`

	n, err := sitter.ParseCtx(context.Background(), []byte(code), jsdoc.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(document (description) (tag (tag_name) (type) (identifier) (description)) (tag (tag_name) (type) (identifier) (description)) (tag (tag_name) (type) (description)))",
		n.String(),
	)
}
