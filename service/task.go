package service

import (
	"encoding/json"
	"fmt"
	"github/com/kakanghosh/jarvis/model"
	"github/com/kakanghosh/jarvis/utils"
	"io/ioutil"
	"os"
	"strconv"
)

func GetTasklist() ([]model.TaskFlags, error) {
	taskList := make([]model.TaskFlags, 0)
	fileData, err := ioutil.ReadFile(utils.TasksFileLocation())
	if err != nil {
		return taskList, err
	}
	err = json.Unmarshal(fileData, &taskList)
	return taskList, err
}

func GetTaskByName(taskName string) *model.TaskFlags {
	taskList, _ := GetTasklist()
	for index, task := range taskList {
		if task.Name == taskName {
			task.Serial = index + 1
			return &task
		}
	}
	return nil
}

func GetTaskBySerial(serial int) *model.TaskFlags {
	taskList, _ := GetTasklist()
	if serial >= 1 && len(taskList) >= serial {
		taskList[serial-1].Serial = serial
		return &taskList[serial-1]
	}
	return nil
}

func GetTaskByNameOrSerial(identifier string) *model.TaskFlags {
	taskList, _ := GetTasklist()
	for _, task := range taskList {
		if task.Name == identifier {
			return &task
		}
	}
	taskNo, err := strconv.Atoi(identifier)
	if err != nil {
		return nil
	}
	if taskNo >= 1 && len(taskList) >= taskNo {
		taskList[taskNo-1].Serial = taskNo
		return &taskList[taskNo-1]
	}
	return nil
}

func RemoveTask(task *model.TaskFlags) error {
	taskList, _ := GetTasklist()
	newTaskList := make([]model.TaskFlags, 0)
	for _, eachTask := range taskList {
		if eachTask.Name != task.Name {
			newTaskList = append(newTaskList, eachTask)
		}
	}
	return saveUpdatedTaskList(newTaskList)
}

func AddTask(task *model.TaskFlags) error {
	taskList, _ := GetTasklist()
	taskList = append(taskList, *task)
	return saveUpdatedTaskList(taskList)
}

func UpdateTask(updateTaskFlags *model.UpdateTaskFlags) (*model.TaskFlags, error) {
	taskList, _ := GetTasklist()

	foundedIndex := updateTaskFlags.Serial - 1

	if len(updateTaskFlags.Name) > 0 {
		if taskWithUpdatedName := GetTaskByName(updateTaskFlags.Name); taskWithUpdatedName != nil {
			if taskWithUpdatedName.Serial != updateTaskFlags.Serial {
				return nil, fmt.Errorf("duplicate name [%s]", updateTaskFlags.Name)
			}
		}
		taskList[foundedIndex].Name = updateTaskFlags.Name
	}
	if updateTaskFlags.WorkingDirectory != "-1" {
		taskList[foundedIndex].WorkingDirectory = updateTaskFlags.WorkingDirectory
	}
	if len(updateTaskFlags.Command) > 0 {
		taskList[foundedIndex].Command = updateTaskFlags.Command
	}
	saveUpdatedTaskList(taskList)
	return &taskList[foundedIndex], nil
}

func saveUpdatedTaskList(taskList []model.TaskFlags) error {
	jsonFile, _ := json.MarshalIndent(taskList, "", "  ")
	return os.WriteFile(utils.TasksFileLocation(), jsonFile, 0644)
}
