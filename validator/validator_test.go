package validator

import (
	"testing"
)

type Model struct {
	Name string `validator:"notEmpty"`
}

func TestValidator(t *testing.T) {
	m := Model{}
	v, err := IsModelValid(m)
	if v {
		t.Error("Model is not valid!", err)
	}

	m = Model{Name: "TestName"}
	v, err = IsModelValid(m)
	if !v {
		t.Error("Model should be valid!", err)
	}
}
