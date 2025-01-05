package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"stroconv"
)

var tasks = []string{}

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Complete task")
		fmt.Println("4. Exit")
		fmt.Println("Enter option: ")

		option, := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option{
		case"1":
			fmt.Print("Enter task:")
			task, := reader.ReadString('\n')
			task = string.TrimSpace(task)
			tasks = append(tasks, task)
			fmt.Println("Task added.")
		case"2":
			fmt.Println("Tasks:")
			for i, task := range tasks{
				fmt.Printf("%d. %s\n", i+1, task)
			}
		case"3":
			fmt.Print("Enter task number to complete: ")
			numStr, _:= reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			num, err := strconv.Atoi(numStr)
			if err == nil && num >= 1 && num <= len(tasks){
				fmt.Printf("Task %d completed.\n" num)
				tasks[num-1] = tasks[num-1] + " [completed]"
			} else{
				fmt.Println("Invalid task number. ")
			}
		case"4":
			fmt.Print("Enter task number to delete: ")
			numStr, _:= reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			num, err := strconv.Atio(numStr)
			if err == nil && num >= 0 && num <= len(tasks){
				tasks = append(tasks[:num-1], tasks[num:]...)
				fmt.Printf("Task %d deleted.\n", num)
			} else{
				fmt.Println("Invalid task number. ")
			}
		case"5":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid optin.")
		}
	}
}