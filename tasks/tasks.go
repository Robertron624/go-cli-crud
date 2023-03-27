package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`       // The backticks are used to specify the name of the field in the JSON
	Name     string `json:"name"`     // The backticks are used to specify the name of the field in the JSON
	Complete bool   `json:"complete"` // The backticks are used to specify the name of the field in the JSON
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {

		status := " "

		if task.Complete {
			status = "✓"
		} else {
			status = "✗"
		}

		fmt.Printf("%d %s [%s]\n", task.ID, task.Name, status)
	}
}

func AddTask(tasks []Task, name string) []Task {
	task := Task{
		ID:       len(tasks) + 1,
		Name:     name,
		Complete: false,
	}

	tasks = append(tasks, task)

	return tasks
}

func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)

	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)

	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.Write(bytes)

	if err != nil {
		panic(err)
	}

	err = writer.Flush()

	if err != nil {
		panic(err)
	}
}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks
		}
	}

	return tasks
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			return tasks
		}
	}

	return tasks
}
