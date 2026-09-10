package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var path string
	fmt.Println("Введите путь к файлу для вывода содержимого построчно(txt):")
	fmt.Scanln(&path)
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
