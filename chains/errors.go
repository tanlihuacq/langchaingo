package chains

import "errors"

var (
	// ErrInvalidInputValues is returned if the input values to a chain is invalid.
	ErrInvalidInputValues = errors.New("invalid input values")
	// ErrMissingInputValues is returned when some expected input values keys to
	// a chain is missing.
	ErrMissingInputValues = errors.New("missing key in input values")
	// ErrInputValuesWrongType is returned if an input value to a chain is of
	// wrong type.
	ErrInputValuesWrongType = errors.New("input key is of wrong type")
	// ErrInvalidOutputValues is returned when expected output keys to a chain does
	// not match the actual keys in the return output values map.
	ErrInvalidOutputValues = errors.New("missing key in output values")

	// ErrMultipleInputsInRun is returned in the run function if the chain expects
	// more then one input values.
	ErrMultipleInputsInRun = errors.New("run not supported in chain with more then one expected input")
	// ErrMultipleOutputsInRun is returned in the run function if the chain expects
	// more then one output values.
	ErrMultipleOutputsInRun = errors.New("run not supported in chain with more then one expected output")
	// ErrMultipleOutputsInRun is returned in the run function if the chain returns
	// a value that is not a string.
	ErrWrongOutputTypeInRun = errors.New("run not supported in chain that returns value that is not string")

	// ErrOutputNotStringInPredict is returned if the output parser in the llm chain
	// returns a value that is not a string in the llm chain.
	ErrOutputNotStringInPredict = errors.New("predict is not supported with a llm chain that does not return a string")
)
