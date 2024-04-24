package evaluator

import (
	"sanevillain/go-interpreter/lexer"
	"sanevillain/go-interpreter/object"
	"sanevillain/go-interpreter/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}
	for _, tC := range testCases {
		evaluated := testEval(tC.input)
		testIntegerObject(t, evaluated, tC.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
	}
	for _, tC := range testCases {
		evaluated := testEval(tC.input)
		testBooleanObject(t, evaluated, tC.expected)
	}
}

func TestBangOperator(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!false", false},
		{"!5", false},
		{"!!5", true},
	}
	for _, tC := range testCases {
		evaluated := testEval(tC.input)
		testBooleanObject(t, evaluated, tC.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d want=%d", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t want=%t", result.Value, expected)
		return false
	}

	return true
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}
