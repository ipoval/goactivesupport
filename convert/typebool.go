package convert

import (
	"strings"
)

// Inspired by ActiveRecord::Type::Boolean.new.cast("true")
// to convert form values from checkbox to the convenient type for an attribute of model
type TypeBoolean struct {
	ValueAsStr string
}

func (t *TypeBoolean) Cast() bool {
	if len(t.ValueAsStr) == 0 {
		return false
	}

	if strings.EqualFold(t.ValueAsStr, "0") ||
		strings.EqualFold(t.ValueAsStr, "false") ||
		strings.EqualFold(t.ValueAsStr, "f") {
		return false
	}

	return true
}
