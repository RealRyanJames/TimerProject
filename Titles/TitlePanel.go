package titles

import (
	"fmt"
	"strings"
)

type UserQuestion struct {
	IsUIShowing bool
}

type LabelMesh struct {
	Label        string
	isUppercase  func() bool
	toLowercase  func() string
	GetTitleJson func() string
}

func (label LabelMesh) UpperTextLabel(str string) string {

	label.isUppercase = func() bool {
		return true
	}

	if label.isUppercase() {

		return fmt.Sprintf("%s", strings.ToUpper(label.GetTitleJson()))
	}

	label.toLowercase = func() string {
		return fmt.Sprintf("%s", strings.ToLower(label.toLowercase()))
	}

	if label.Label == "Uppercase" {
		return fmt.Sprintf("%s", strings.ToUpper(label.Label))
	}

	label.Label = str
	return label.Label
}
