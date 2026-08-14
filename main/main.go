package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/JxSam/go-bot/pet-project/learning/function"
	hashtable "github.com/JxSam/go-bot/pet-project/learning/hash_table"
	xlsxconverttomap "github.com/JxSam/go-bot/pet-project/xlsx_convert_to_map"
)

func main() {
	var action string
	var projects = map[int]string{
		1: "csv_convert_to_map",
		2: "learning_projects",
	}
	var learningProjects = map[int]string{
		1: "function | even_noeven",
		2: "hash_table | money converter",
	}
	fmt.Println("Pet-projects by JxSam aka N1kS. Select action:")
	printMapTable(projects)
	fmt.Scan(&action)
	switch action {
	case "1":
		xlsxconverttomap.Convert()
	case "2":
		fmt.Println("Learnign projectss:")
		printMapTable(learningProjects)
		fmt.Scan(&action)
		switch action {
		case "1":
			function.Even_noeven()
		case "2":
			hashtable.Convert()
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
