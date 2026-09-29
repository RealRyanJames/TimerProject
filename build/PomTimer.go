package main

import (
	datesoptions "PointerTimer/DatesOptions"
	parsing "PointerTimer/Parsing"
	uistyles "PointerTimer/UIStyles"
	userinfostate "PointerTimer/UserInfoState"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type JsonDataMesh struct {
	Title       string
	VERSION     string
	ToLowerCase string
	isUpperCase bool
}

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

	content, err := os.ReadFile("../Configs/ConfigParams.json")

	if err != nil {
		fmt.Println(err)
	}

	var jsonLoadsData JsonDataMesh
	err = json.Unmarshal(content, &jsonLoadsData)

	if err != nil {
		fmt.Println(err)
	}

	parsedData := parsing.ParseingMeshType{
		IsDataShowing: func() bool {
			return true
		},

		GetVersionData: func() string {
			return jsonLoadsData.VERSION
		},

		GetParsedData: func() string {

			return strings.ToUpper(jsonLoadsData.Title)
		},
	}

	if parsedData.GetParsedDataJson() {

		fmt.Println("Welcome to:", parsedData.GetParsedDataData())

		fmt.Println("Current Version:", parsedData.GetVersionData())
	}

	fmt.Printf("%s", "\n")

	fmt.Println("RUNNING APPLICATION")
	uistyles.AppendLinesUI()
	fmt.Printf("%s", "\n")

	dateNow := datesoptions.DateOptions{}
	fmt.Printf("%d/%d/%d\n", dateNow.Get().Date.Month(), dateNow.Get().Date.Day(), dateNow.Get().Date.Year())

	userPrompt := userinfostate.UserQuestion{
		IsUIShowing: true,
	}

	sQuestion := ""
	questionByUser := userinfostate.ReadUserQuestion{
		Str: sQuestion,
	}

	userPrompt.GetUser()

	fmt.Scanln(&questionByUser.Str)

	if questionByUser.GetReadSync() == "20M" {

		sys_func := SystemMessageMesh{
			Command: strings.ToLower("PAUSE"),
		}

		time_current := 0
		for time_current > -1 {

			if time_current == 20000 {

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
	}

	if questionByUser.GetReadSync() == "30M" {

		sys_func := SystemMessageMesh{
			Command: strings.ToLower("PAUSE"),
		}

		time_current := 0
		for time_current > -1 {

			if time_current == 30000 {

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
	}

	if questionByUser.GetReadSync() == "5M" {

		sys_func := SystemMessageMesh{
			Command: strings.ToLower("PAUSE"),
		}

		time_current := 0
		for time_current > -1 {

			if time_current == 5000 {

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
	}

	if questionByUser.GetReadSync() == "/O" {
		fmt.Println("1  Hour: H1")
		fmt.Println("2  Hours: H2")
		fmt.Println("5  Minutes: 5M")
		fmt.Println("20 Minutes: 20M")
		fmt.Println("30 Minutes: 30M")

	} else if questionByUser.GetReadSync() == "H1" {

		sys_func := SystemMessageMesh{
			Command: strings.ToLower("PAUSE"),
		}

		time_current := 0
		for time_current > -1 {

			if time_current == 60000 {

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
	} else if questionByUser.GetReadSync() == "H2" {

		sys_func := SystemMessageMesh{
			Command: strings.ToLower("PAUSE"),
		}

		time_current := 0
		for time_current > -1 {

			if time_current == 60*2 {

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

	}

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
