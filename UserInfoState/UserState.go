package userinfostate

import (
	"fmt"
	"strings"
)

type UserQuestion struct {
	IsUIShowing bool
}

type LabelMesh struct {
	Label string
}

func (label LabelMesh) UpperTextLabel(str string) string {

	label.Label = str
	return label.Label
}

func (userInfoUI UserQuestion) GetUser() {
	if userInfoUI.IsUIShowing {
		labelUser := LabelMesh{

			Label: fmt.Sprintf("%s", strings.ToUpper("=> Please Enter Answer: ")),
		}

		fmt.Println(labelUser.UpperTextLabel(labelUser.Label))
	}
}

type ReadUserQuestion struct {
	LoggerRunning func()
	Str           string
}

func (loggerRunning ReadUserQuestion) GetReadSync() string {

	return loggerRunning.Str
}
