package gucassowary

// Multiply defines the multiple operator.
var Multiply multiply

// Multiply defines the handle for the multiplication operator.
type multiply struct{}

// Do resolves a multiplication of number type values.
func (m multiply) Do(left, right Variable, res Variable) {
	leftValue := left.Value()
	rightValue := right.Value()

	switch leftValue.(type) {
	case float32, float64:
		if rv, ok := rightValue.(float64); ok {
			res.Set(rv * leftValue.(float64))
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(rv * leftValue.(float64))
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(float64(rv) * leftValue.(float64))
			return
		}

	case int, int32, int64, int8, int16:
		lv := leftValue.(int)

		if rv, ok := rightValue.(float64); ok {
			res.Set(rv * float64(lv))
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(rv * float64(lv))
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(lv * rv)
			return
		}
	}
}

//==============================================================================

// Add provides the addition operator.
var Add addition

// addition defines the handle for the Addition operator.
type addition struct{}

// Do resolves a multiplication of number type values.
func (a addition) Do(left, right Variable, res Variable) {
	leftValue := left.Value()
	rightValue := right.Value()

	switch leftValue.(type) {
	case float32, float64:
		if rv, ok := rightValue.(float64); ok {
			res.Set(rv + leftValue.(float64))
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(rv + leftValue.(float64))
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(float64(rv) + leftValue.(float64))
			return
		}

	case int, int32, int64, int8, int16:
		lv := leftValue.(int)

		if rv, ok := rightValue.(float64); ok {
			res.Set(rv + float64(lv))
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(rv + float64(lv))
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(lv + rv)
			return
		}
	}
}

//==============================================================================

// Subtract provides the Subtraction operator.
var Subtract subtraction

// Subtraction defines the handle for the Addition operator.
type subtraction struct{}

// Do resolves a multiplication of number type values.
func (s subtraction) Do(left, right Variable, res Variable) {
	leftValue := left.Value()
	rightValue := right.Value()

	switch leftValue.(type) {
	case float32, float64:
		if rv, ok := rightValue.(float64); ok {
			res.Set(leftValue.(float64) - rv)
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(leftValue.(float64) - rv)
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(leftValue.(float64) + float64(rv))
			return
		}

	case int, int32, int64, int8, int16:
		lv := leftValue.(int)

		if rv, ok := rightValue.(float64); ok {
			res.Set(float64(lv) - rv)
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(float64(lv) - rv)
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(lv - rv)
			return
		}
	}
}

//==============================================================================

// Divide provides the divide operator.
var Divide division

// division defines the handle for the Addition operator.
type division struct{}

// Do resolves a multiplication of number type values.
func (d division) Do(left, right Variable, res Variable) {
	leftValue := left.Value()
	rightValue := right.Value()

	switch leftValue.(type) {
	case float32, float64:
		if rv, ok := rightValue.(float64); ok {
			res.Set(leftValue.(float64) / rv)
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(leftValue.(float64) / rv)
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(leftValue.(float64) / float64(rv))
			return
		}

	case int, int32, int64, int8, int16:
		lv := leftValue.(int)

		if rv, ok := rightValue.(float64); ok {
			res.Set(float64(lv) / rv)
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(float64(lv) / rv)
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(float64(lv) / float64(rv))
			return
		}
	}
}

//==============================================================================

// Mod provides the % operator.
var Mod modulus

// division defines the handle for the Addition operator.
type modulus struct{}

// Do resolves a multiplication of number type values.
func (d modulus) Do(left, right Variable, res Variable) {
	leftValue := left.Value()
	rightValue := right.Value()

	switch leftValue.(type) {
	case float32, float64:
		if rv, ok := rightValue.(float64); ok {
			res.Set(leftValue.(float64) % rv)
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(leftValue.(float64) % rv)
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(leftValue.(float64) % float64(rv))
			return
		}

	case int, int32, int64, int8, int16:
		lv := leftValue.(int)

		if rv, ok := rightValue.(float64); ok {
			res.Set(float64(lv) % rv)
			return
		}

		if rv, ok := rightValue.(float32); ok {
			res.Set(float64(lv) % rv)
			return
		}

		if rv, ok := rightValue.(int); ok {
			res.Set(lv % rv)
			return
		}
	}
}

//==============================================================================
