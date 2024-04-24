package evaluator

import (
	"sanevillain/go-interpreter/ast"
	"sanevillain/go-interpreter/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}
