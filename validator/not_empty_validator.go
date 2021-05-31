package validator

import (
	"errors"
	"reflect"
	"strings"
)

func IsNotEmpty(v reflect.Value) (bool, error) {
	if len(strings.Trim(v.String(), "\n\t\r ")) > 0 {
		return true, nil
	}

	return false, errors.New("string is empty")
}
