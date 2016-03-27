package gucassowary

//==============================================================================

// Expressor defines a actionable expression.
type Expressor interface {
	Do(Variable)
}

// Expression defines the interface for different operations to be
// performed by sets of expressions. Its signature allows producing inner
// chains of internal expressions.
type Expression interface {
	Expressor
	Right() Variable
	Left() Variable
	Operation() Operator
}

// Operator provides an interface that defines an operation to be performed
// against two varables.
type Operator interface {
	Do(left, right Variable, result Variable)
}

//==============================================================================

// Expr defines a concrete implementation of the Expression interface.
type Expr struct {
	left  Variable
	right Variable
	op    Operator
}

// Exprs returns a new Expression struct
func Exprs(op Operator, left, right Variable) Expression {
	expr := Expr{left: left, right: right, op: op}
	return &expr
}

// Right returns the right variable of this expression.
func (e Expr) Right() Variable {
	return e.right
}

// Left returns the left variable of this expression.
func (e Expr) Left() Variable {
	return e.left
}

// Operation returns the operation for this expression.
func (e Expr) Operation() Operator {
	return e.op
}

// Do performs the operation on the left and right variable using the giving
// operation.
func (e Expr) Do(res Variable) {
	if e.op != nil {
		return e.op.Do(e.left, e.right, res)
	}

	return nil
}

//==============================================================================
