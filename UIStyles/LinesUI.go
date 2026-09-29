package uistyles

import (
	"fmt"
	"strings"
)

type LineUI struct {
	linesUIStyle string
	posX         float64
	posY         float64
}

type Calculations struct {
	GetPosX func() float64
}

func MultiplyPosLines(x float64, y float64) float64 {
	return float64((x * y) / 2)
}

func (PosCalc Calculations) GetPositions() float64 {

	return MultiplyPosLines(PosCalc.GetPosX(), 2.0)
}

type VersionApplication struct {
	Name    string
	IsOwner bool
}

func (versionId VersionApplication) GetId() string {
	if versionId.IsOwner {
		return "Welcome To: " + strings.ToUpper(versionId.Name)
	}

	return "Welcome To: " + strings.ToUpper(versionId.Name)
}

func AppendLinesUI() {

	posLines := LineUI{
		linesUIStyle: "-",
		posX:         40.0,
		posY:         2.0,
	}

	calc := Calculations{
		GetPosX: func() float64 {
			return 40
		},
	}

	for i := 0; i < int(calc.GetPositions()); i++ {

		fmt.Print(posLines.linesUIStyle)
	}

}
