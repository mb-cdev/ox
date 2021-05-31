package validator

import (
	"errors"
	"reflect"
	"sync"
)

type ValidatorFunc func(reflect.Value) (bool, error)

var errNilModel = errors.New("the model is nil")
var validators map[string]ValidatorFunc

func init() {
	var once sync.Once

	once.Do(func() {
		validators = make(map[string]ValidatorFunc)

		validators["notEmpty"] = IsNotEmpty
	})
}

func IsModelValid(model interface{}) (bool, error) {
	typ := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	if typ == nil {
		return false, errNilModel
	}

	for i := 0; i < typ.NumField(); i++ {
		fld := typ.Field(i)
		t := fld.Tag.Get("validator")
		if t == "" {
			continue
		}
		validatorFunc, ok := validators[t]
		if !ok {
			continue
		}

		if ok, err := validatorFunc(v.Field(i)); !ok {
			return false, err
		}
	}

	return true, nil
}
