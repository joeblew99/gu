package gucassowary

//==============================================================================

// VariableType defines the types of variables used within the constraint.
type VariableType int

const (
	// NoType defines a notype variable and should not be used
	NoType VariableType = iota
	// Slack represents slack variable types
	Slack
	// Dummy represents Dummy variable types
	Dummy
	// Objective represents objective variable types
	Objective
	// Variable represents ordinary variable types
	Variable
)

//==============================================================================

// Variable defines an interface for a variable type.
type Variable interface {
	Value() interface{}
	Type() VariableType
	Name() string
	Set(interface{}) bool
}

//==============================================================================

// varc provides a based variable type struct.
type varc struct {
	vt    VariableType
	value interface{}
	name  string
}

// Value returns the internal value of the variable.
func (v *varc) Value() interface{} {
	return v.value
}

// Type returns the variable type for the given variable.
func (v *varc) Type() VariableType {
	return v.vt
}

// Name returns the name of the given variable.
func (v *varc) Name() string {
	return v.name
}

//==============================================================================

// Vars defines a concrete variable which can be change
// using the set value.
type Vars struct {
	*varc
}

// Var returns a new variable instance.
func Var(vt VariableType, name string, vm interface{}) Variable {
	vc := varc{vt: vt, name: name, value: vm}
	return &Vars{varc: &vc}
}

// Set changes the internal value held by the variable,
// returning a true/false if it was successful.
func (v *Vars) Set(val interface{}) bool {
	v.value = val
	return true
}

//==============================================================================

// Consts defines a constant variable, which ones set can
// not be changed/altered.
type Consts struct {
	*varc
}

// Const returns a new constant variable set with the given value.
func Const(vt VariableType, name string, m interface{}) Variable {
	vc := varc{vt: vt, name: name, value: vm}
	return &Consts{varc: &vc}
}

// Set changes the internal value held by the variable,
// returning a true/false if it was successful.
func (c *Const) Set(_ interface{}) bool {
	return false
}
