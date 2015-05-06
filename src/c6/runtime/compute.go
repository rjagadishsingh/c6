package runtime

import "fmt"
import "c6/ast"

/*
Value
*/
type ComputeFunction func(a ast.Value, b ast.Value) ast.Value

const ValueTypeNum = 7

var computableMatrix [ValueTypeNum][ValueTypeNum]bool = [ValueTypeNum][ValueTypeNum]bool{
	/* NumberValue */
	[ValueTypeNum]bool{false, false, false, false, false, false, false},

	/* LengthValue */
	[ValueTypeNum]bool{false, false, false, false, false, false, false},

	/* HexColorValue */
	[ValueTypeNum]bool{false, false, false, false, false, false, false},

	/* RGBAColorValue */
	[ValueTypeNum]bool{false, false, false, false, false, false, false},

	/* RGBColorValue */
	[ValueTypeNum]bool{false, false, false, false, false, false, false},
}

/**
Each row: [5]ComputeFunction{ NumberValue, LengthValue, HexColorValue, RGBAColorValue, RGBColorValue }
*/
var computeFunctionMatrix [5][5]ComputeFunction = [5][5]ComputeFunction{

	/* NumberValue */
	[5]ComputeFunction{nil, nil, nil, nil, nil},

	/* LengthValue */
	[5]ComputeFunction{nil, nil, nil, nil, nil},

	/* HexColorValue */
	[5]ComputeFunction{nil, nil, nil, nil, nil},

	/* RGBAColorValue */
	[5]ComputeFunction{nil, nil, nil, nil, nil},

	/* RGBColorValue */
	[5]ComputeFunction{nil, nil, nil, nil, nil},
}

func Compute(op *ast.Op, a ast.Value, b ast.Value) ast.Value {
	switch op.Type {
	case ast.T_PLUS:
		switch ta := a.(type) {
		case *ast.Number:
			switch tb := b.(type) {
			case *ast.Number:
				return NumberAddNumber(ta, tb)
			case *ast.HexColor:
				return HexColorAddNumber(tb, ta)
			}
		case *ast.Length:
			switch tb := b.(type) {
			case *ast.Length:
				return LengthAddLength(ta, tb)
			}
		case *ast.HexColor:
			switch tb := b.(type) {
			case *ast.Number:
				return HexColorAddNumber(ta, tb)
			}
		case *ast.RGBColor:
			switch tb := b.(type) {
			case *ast.Number:
				return RGBColorAddNumber(ta, tb)
			}
		case *ast.RGBAColor:
			switch tb := b.(type) {
			case *ast.Number:
				return RGBAColorAddNumber(ta, tb)
			}
		}
	case ast.T_MINUS:
		switch ta := a.(type) {

		case *ast.Number:
			switch tb := b.(type) {
			case *ast.Number:
				return NumberSubNumber(ta, tb)
			}

		case *ast.Length:
			switch tb := b.(type) {
			case *ast.Length:
				val := LengthSubLength(ta, tb)
				fmt.Printf("Substracted value: %+v\n", val)
				return val
			}

		case *ast.HexColor:
			switch tb := b.(type) {
			case *ast.Number:
				return HexColorSubNumber(ta, tb)
			}

		case *ast.RGBColor:
			switch tb := b.(type) {
			case *ast.Number:
				return RGBColorSubNumber(ta, tb)
			}

		case *ast.RGBAColor:
			switch tb := b.(type) {
			case *ast.Number:
				return RGBAColorSubNumber(ta, tb)
			}
		}
	case ast.T_MUL:
		switch ta := a.(type) {

		case *ast.Length:
			switch tb := b.(type) {
			case *ast.Length:
				return LengthMulLength(ta, tb)
			case *ast.Number:
				return LengthMulNumber(ta, tb)
			}

		case *ast.RGBColor:
			switch tb := b.(type) {
			case *ast.Number:
				return RGBColorMulNumber(ta, tb)
			}

		case *ast.RGBAColor:
			switch tb := b.(type) {
			case *ast.Number:
				return RGBAColorMulNumber(ta, tb)
			}
		}
	}
	return nil
}

func EvaluateBinaryExpression(expr *ast.BinaryExpression, symTable *SymTable) ast.Value {
	if expr.IsCssSlash() {
		// return string object without quote
		return ast.NewString(0, expr.Left.(*ast.Length).String()+"/"+expr.Right.(*ast.Length).String(), nil)
	}

	var lval ast.Value = nil
	var rval ast.Value = nil

	switch expr := expr.Left.(type) {
	case *ast.BinaryExpression:
		lval = EvaluateBinaryExpression(expr, symTable)
	case *ast.UnaryExpression:
		lval = EvaluateUnaryExpression(expr, symTable)
	case *ast.Number, *ast.Length, *ast.HexColor:
		lval = ast.Value(expr)
	}
	switch expr := expr.Right.(type) {
	case *ast.UnaryExpression:
		rval = EvaluateUnaryExpression(expr, symTable)
	case *ast.BinaryExpression:
		rval = EvaluateBinaryExpression(expr, symTable)
	case *ast.Number, *ast.Length, *ast.HexColor:
		rval = ast.Value(expr)
	}
	if lval != nil && rval != nil {
		return Compute(expr.Op, lval, rval)
	}
	return nil
}

func EvaluateUnaryExpression(expr *ast.UnaryExpression, symTable *SymTable) ast.Value {
	var val ast.Value = nil
	if bexpr, ok := expr.Expr.(*ast.BinaryExpression); ok {
		val = EvaluateBinaryExpression(bexpr, symTable)
	} else if uexpr, ok := expr.Expr.(*ast.UnaryExpression); ok {
		val = EvaluateUnaryExpression(uexpr, symTable)
	}

	// negative value
	if expr.Op.Type == ast.T_DIV {
		switch n := val.(type) {
		case *ast.Number:
			n.Value = -n.Value
		case *ast.Length:
			n.Value = -n.Value
		}
	}
	return val
}