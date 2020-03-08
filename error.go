package vipervar

import "fmt"

// ErrInvalidDelimiter returns when a start or end delimiter is
// empty or contains an invalid character
type ErrInvalidDelimiter struct {
	delim  string
	reason string
}

func (e ErrInvalidDelimiter) Error() string {
	return fmt.Sprintf("Delimiter `%s` invalid: %s", e.delim, e.reason)
}

// ErrNonExistentVariableReference returns is a variable
// references a non-existant value in the configured Viper
type ErrNonExistentVariableReference struct {
	variable string
}

func (e ErrNonExistentVariableReference) Error() string {
	return fmt.Sprintf(
		"Variable `%s` references a non-existant value",
		e.variable,
	)
}

// ErrUnsupportedVariableType returns if a variable resolves to
// a type where no assumptions can be made for assignment, such
// as maps and slices
type ErrUnsupportedVariableType struct {
	variable string
	vartype  interface{}
}

func (e ErrUnsupportedVariableType) Error() string {
	return fmt.Sprintf(
		"Variable `%s` references a variable with an unsupported type: %T",
		e.variable,
		e.vartype,
	)
}

// ErrNonExistentKey returns when a specified key does not exist
type ErrNonExistentKey struct {
	key string
}

func (e ErrNonExistentKey) Error() string {
	return fmt.Sprintf(
		"Key `%s` does not exist",
		e.key,
	)
}
