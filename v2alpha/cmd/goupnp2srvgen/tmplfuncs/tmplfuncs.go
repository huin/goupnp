// Package tmplfuncs contains functions for injection into templates.
package tmplfuncs

import (
	"errors"
	"fmt"
)

// Args accepts pairs of string names and any values and constructs a map from them.
// This is to help passing multiple arguments to a template from within a template.
func Args(args ...any) (map[string]any, error) {
	if len(args)%2 != 0 {
		return nil, errors.New("args must have an even number of arguments")
	}
	res := make(map[string]any, len(args)/2)

	for i := 0; i < len(args); i += 2 {
		name, ok := args[i].(string)
		if !ok {
			return nil, fmt.Errorf("argument %d: want string, got %T", i, args[i])
		}
		value := args[i+1]
		if _, exists := res[name]; exists {
			return nil, fmt.Errorf("argument name %q occurs more than once", name)
		}
		res[name] = value
	}

	return res, nil
}
