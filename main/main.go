package main

import (
	"fmt"
	"sort"
	"strings"

	moneyconverter "github.com/JxSam/go-bot/pet-project/money_converter"
)

func main() {
	var action string
	var projects = map[int]string{
		1: "money_converter",
	}
	fmt.Println("Pet-projects by JxSam aka N1kS. Select action:")
	printMapTable(projects)
	fmt.Scan(&action)
	switch action {
	case "1":
		moneyconverter.Convert()
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
