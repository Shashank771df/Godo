package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Todo struct {
	ID          int
	Description string
	Done        bool
}

var todos []Todo

var cyan = color.New(color.FgCyan).SprintFunc()

const fileName = "task.txt"
const seprator = "--"

var idCounter int

func main() {

	idCounter = len(todos) + 1
	var choice string

	logo()
	loadTasksFormFile()
	for {
		showMenu()
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("Add a task")
			addTask()
		case "2":
			fmt.Println("View all tasks")
			listAllTasks()

		case "3":
			fmt.Println("Delete a task")
			deleteTask()

		case "4":
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}

}

func logo() {
	fmt.Printf("%s\n", cyan(`

	███████╗██╗  ██╗ █████╗ ███████╗██╗  ██╗ █████╗ ███╗   ██╗██╗  ██╗███████╗███████╗ ██╗██████╗ ███████╗
	██╔════╝██║  ██║██╔══██╗██╔════╝██║  ██║██╔══██╗████╗  ██║██║ ██╔╝╚════██║╚════██║███║██╔══██╗██╔════╝
	███████╗███████║███████║███████╗███████║███████║██╔██╗ ██║█████╔╝     ██╔╝    ██╔╝╚██║██║  ██║█████╗  
	╚════██║██╔══██║██╔══██║╚════██║██╔══██║██╔══██║██║╚██╗██║██╔═██╗    ██╔╝    ██╔╝  ██║██║  ██║██╔══╝  
	███████║██║  ██║██║  ██║███████║██║  ██║██║  ██║██║ ╚████║██║  ██╗   ██║     ██║   ██║██████╔╝██║     
	╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝   ╚═╝     ╚═╝   ╚═╝╚═════╝ ╚═╝     
	Project: Godo - CLI Task Manager
	Author: Shashank Chaudhary
	Version: 1.0.0																									  	   
    `))
	fmt.Println("Welcome to the task manager!")
}

func showMenu() {
	fmt.Println("1. Add a task")
	fmt.Println("2. View all tasks")
	fmt.Println("3. Delete a task")
	fmt.Println("4. Exit")
}

func loadTasksFormFile() {
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		if strings.Contains(fileErr.Error(), "The system cannot find the file specified") {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("\n creating db for you ....")
			time.Sleep(200 * time.Millisecond)
			fmt.Println("\n db created successfully")
			file, fileErr = os.Create(fileName)
			if fileErr != nil {
				fmt.Println("Error creating file")
			}
		} else {
			fmt.Println("Error opening file", fileErr)
		}
	}
	defer file.Close()

	//content := bufio.NewReader(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, seprator)

		if len(fields) != 3 {
			fmt.Println("Invalid data in file:", line)
			continue
		}

		id, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Invalid ID in file:", fields[0])
			continue
		}

		status, err := strconv.ParseBool(fields[2])
		if err != nil {
			fmt.Println("Invalid status in file:", fields[2])
			continue
		}
		todo := Todo{
			ID:          id,
			Description: fields[1],
			Done:        status,
		}
		todos = append(todos, todo)
	}
}

func saveTasksToFile() {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, todo := range todos {
		line := fmt.Sprintf("%d"+seprator+"%s"+seprator+"%t\n", todo.ID, todo.Description, todo.Done)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func listAllTasks() {
	for _, todo := range todos {
		status := "Not-Done"
		if todo.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", todo.ID, todo.Description, status)
	}
}

func addTask() {
	fmt.Print("Enter task description: ")
	var description string
	fmt.Scanln(&description)

	if len(todos) > 0 {
		idCounter++
	}
	todo := Todo{
		ID:          idCounter,
		Description: description,
		Done:        false,
	}
	todos = append(todos, todo)
	saveTasksToFile()

}

func deleteTask() {
	fmt.Print("Enter task ID to delete: ")
	var id int
	fmt.Scanln(&id)

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	saveTasksToFile()

}
