package hel

import (
	"github.com/hashicorp/hel/ast"
)

// FixedValueTransform transforms an AST to return a fixed value for
// all interpolations. i.e. you can make "hello ${anything}" always
// turn into "hello foo".
//
// The primary use case for this is for config validations where you can
// verify that interpolations result in a certain type of string.
func FixedValueTransform(root ast.Node, Value *ast.LiteralNode) ast.Node {
	// We visit the nodes in top-down order
	result := root
	switch n := result.(type) {
	case *ast.Concat:
		for i, v := range n.Exprs {
			n.Exprs[i] = FixedValueTransform(v, Value)
		}
	case *ast.LiteralNode:
		// We keep it as-is
	default:
		// Anything else we replace
		result = Value
	}

	return result
}