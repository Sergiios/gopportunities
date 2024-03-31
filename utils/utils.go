package utils

import "fmt"

func ParamRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}
