package evaluator

import (
	"sanevillain/go-interpreter/ast"
	"sanevillain/go-interpreter/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	return nil
}

func evalProgram(node *ast.Program, env *object.Environment) object.Object {
	return nil
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	return nil
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	return nil
}

func isTruthy(obj object.Object) bool {
	return false
}

func evalExpressions(exprs []ast.Expression, env *object.Environment) []object.Object {
	return nil
}

func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	return nil
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	return nil
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	return nil
}

func evalIndexExpression(left, index object.Object) object.Object {
	return nil
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	return nil
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	return nil
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	return nil
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	return nil
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	return nil
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	return nil
}

func evalBangOperatorExpression(right object.Object) object.Object {
	return nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	return nil
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	return nil
}

func unwrapReturnValue(obj object.Object) object.Object {
	return nil
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	return nil
}

func isError(obj object.Object) bool {
	return false
}

func newError(format string, a ...any) *object.Error {
	return nil
}
