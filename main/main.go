package main

import (
	"fmt"
	"sort"
	"strings"

	funca "github.com/JxSam/go-bot/pet-project/learning/function"
	hashtable "github.com/JxSam/go-bot/pet-project/learning/hash_table"
	xlsxconverttomap "github.com/JxSam/go-bot/pet-project/xlsx_convert_to_map"
)

func main() {
	projects := map[int]string{
		1: "csv_convert_to_map",
		2: "learning_projects | ...",
	}
	learningProjects := map[int]string{
		1: "function | ...",
		2: "hash_table | ...",
		0: "<Back",
	}
	function := map[int]string{
		1: "even_noeven",
		2: "calculator",
		3: "filter",
		0: "<Back",
	}
	hash_table := map[int]string{
		1: "money_converter",
		0: "<Back",
	}

	var level int // 0 = главное меню, 1 = learning projects, 2 = function
	var action string

	for {
		switch level {
		case 0:
			fmt.Println("Pet-projects by JxSam aka N1kS. Select action:")
			printMapTable(projects)
			fmt.Scan(&action)
			switch action {
			case "1":
				xlsxconverttomap.Convert()
			case "2":
				level = 1 // переходим на уровень learningProjects
			default:
				fmt.Println("Invalid option")
			}

		case 1:
			fmt.Println("Learning projects:")
			printMapTable(learningProjects)
			fmt.Scan(&action)
			switch action {
			case "1":
				level = 2 // переходим на уровень function
			case "2":
				level = 3
			case "<Back", "0": // можно реагировать и на текст, и на номер
				level = 0 // назад в главное меню
			default:
				fmt.Println("Invalid option")
			}

		case 2:
			printMapTable(function)
			fmt.Scan(&action)
			switch action {
			case "1":
				funca.Even_noeven()
			case "2":
				funca.Calculator()
			case "3":
				funca.MainFilter()
			case "0", "<Back":
				level = 1 // назад к learningProjects
			default:
				fmt.Println("Invalid option")
			}
		case 3:
			printMapTable(hash_table)
			fmt.Scan(&action)
			switch action {
			case "1":
				hashtable.Convert()
			case "0":
				level = 1
			}
		}
	}
}

func printMapTable(m map[int]string) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Printf("%-10s | %s\n", "ID", "Project")
	fmt.Println(strings.Repeat("-", 45))

	for _, k := range keys {
		fmt.Printf("%-10d | %s\n", k, m[k])
	}
}
