package parsing

type ParseingMeshType struct {
	GetParsedData  func() string
	IsDataShowing  func() bool
	GetVersionData func() string
}

func (parsedData ParseingMeshType) GetParsedDataJson() bool {

	if parsedData.IsDataShowing() {
		return true
	}

	return false
}

func (parsedData ParseingMeshType) GetParsedDataData() string {

	return "TimerPom"

}
