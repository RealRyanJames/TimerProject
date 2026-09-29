package main

import (
	datesoptions "PointerTimer/DatesOptions"
	uistyles "PointerTimer/UIStyles"
	"fmt"
	"strings"
	"time"
)

type MainTaskMesh struct {
	IdInt     int
	isWritten bool
}

type TodoMessageMesh struct {
	errMessageInit string
}

func (todoMesh TodoMessageMesh) GetErrorFromTodo() string {
	return strings.ToUpper(todoMesh.errMessageInit)
}

type funcNamePicker struct {
	nameMessage string
}

func (messagePersonInit funcNamePicker) GetName() string {
	return messagePersonInit.nameMessage
}

type MeshesUser struct {
	outputMessage string
	isUnderMesh   bool
}

func readStringOuput(str string) string {
	strMessage := fmt.Sprintf("%s", str)
	return strMessage
}

func (Mesh MeshesUser) ReadMesh() string {
	if Mesh.isUnderMesh {
		return readStringOuput(strings.ToUpper(Mesh.outputMessage))
	}

	return readStringOuput(strings.ToUpper(Mesh.outputMessage))
}

type SystemMessageMesh struct {
	Command         string
	SystemFuncMeshc func()
}

type dateNow struct {
}

func main() {

	uistyles.AppendLinesUI()
	fmt.Printf("%s", "\n")

	versionUser := uistyles.VersionApplication{
		Name:    "Program User",
		IsOwner: true,
	}

	fmt.Print(versionUser.GetId())
	fmt.Printf("%s", "\n")

	fmt.Println("RUNNING APPLICATION")
	uistyles.AppendLinesUI()
	fmt.Printf("%s", "\n")

	dateNow := datesoptions.DateOptions{}
	fmt.Printf("%d/%d/%d\n", dateNow.Get().Date.Month(), dateNow.Get().Date.Day(), dateNow.Get().Date.Year())

	sys_func := SystemMessageMesh{
		Command: strings.ToLower("PAUSE"),
	}

	sys_func.SystemFuncMeshc = func() {
		fmt.Scanln()
	}

	time_current := 0
	for time_current > -1 {

		if time_current == 25*60 {

			value := fmt.Sprintf("%s", "done")
			time.Sleep(1 * time.Second)

			meshesUpper := MeshesUser{
				isUnderMesh:   true,
				outputMessage: value,
			}
			fmt.Println(strings.ToUpper(meshesUpper.ReadMesh()))
			sys_func.SystemFuncMeshc()
			break
		}

	}

	sys_func.SystemFuncMeshc()
}
