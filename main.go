package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID     int
	Title  string
	Status string
}

var tasks []Task

func main() {
	tasks = []Task{
		{1, "Setup project", "todo"},
		{2, "Write code", "inprogress"},
		{3, "Review", "done"},
	}

	menu()
}

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	runMenu(scanner)
}

func runMenu(scanner *bufio.Scanner) {
	for {
		printMenu()

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {

		case "1":
			fmt.Print("Title: ")
			scanner.Scan()

			id := len(tasks) + 1
			tasks = append(tasks, Task{id, scanner.Text(), "todo"})

			fmt.Println("Task created.")

		case "2":
			fmt.Print("Enter ID: ")
			scanner.Scan()

			id, _ := strconv.Atoi(scanner.Text())

			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Status = "done"
					fmt.Println("Marked done.")
				}
			}

		case "3":
			fmt.Print("Enter ID: ")
			scanner.Scan()

			id, _ := strconv.Atoi(scanner.Text())

			for i := range tasks {
				if tasks[i].ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("Deleted.")
					break
				}
			}

		case "4":
			printTasks(tasks)

		case "5":
			fmt.Println("Программа завершена.")
			return

		default:
			fmt.Println("Введите число от 1 до 5.")
		}
	}
}

func printMenu() {
	fmt.Println("== МЕНЮ ЗАДАЧ ==")
	fmt.Println("1. Создать таску")
	fmt.Println("2. Пометить выполненным")
	fmt.Println("3. Удалить таску")
	fmt.Println("4. Все таски")
	fmt.Println("5. Завершить работу программы")
	fmt.Print("Выберите действие: ")
}

func printTasks(tasks []Task) {
	fmt.Println("+----+----------------------+-------------+")
	fmt.Println("| ID | Title                | Status      |")
	fmt.Println("+----+----------------------+-------------+")

	for _, task := range tasks {
		fmt.Printf("| %2d | %-20s | %-11s |\n",
			task.ID, task.Title, task.Status)
	}

	fmt.Println("+----+----------------------+-------------+")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("=== МЕНЮ ЗАДАЧ ===")
		fmt.Println("1. Создать таску")
		fmt.Println("2. Пометить выполненным")
		fmt.Println("3. Удалить таску")
		fmt.Println("4. Все таски")
		fmt.Println("5. Завершить работу программы")
		fmt.Println("Выберите действие (1-5):")

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Println("Создание задач")
		case "2":
			fmt.Println("Выполнение задачи")
		case "3":
			fmt.Println("Удаление задачи")
		case "4":
			fmt.Println("Вывод всех задач")
		case "5":
			return
		default:
			fmt.Println("Неверный ввод. Введите число от 1 до 5.")
		}
	}
}

func main() {
	menu()
}
