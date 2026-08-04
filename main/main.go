package main

import (
	"fmt"
	"sort"
	"strings"

	easyprojects "github.com/JxSam/go-bot/pet-project/easy_projects"
	moneyconverter "github.com/JxSam/go-bot/pet-project/money_converter"
	xlsxconverttomap "github.com/JxSam/go-bot/pet-project/xlsx_convert_to_map"
)

func main() {
	var action string
	var projects = map[int]string{
		1: "money_converter",
		2: "csv_convert_to_map",
		3: "easy_projects",
	}
	var easy_projects = map[int]string{
		1: "is_even?",
	}
	fmt.Println("Pet-projects by JxSam aka N1kS. Select action:")
	printMapTable(projects)
	fmt.Scan(&action)
	switch action {
	case "1":
		moneyconverter.Convert()
	case "2":
		xlsxconverttomap.Convert()
	case "3":
		fmt.Println("Easy projects:")
		printMapTable(easy_projects)
		fmt.Scan(&action)
		switch action {
		case "1":
			easyprojects.Even_noeven()
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
