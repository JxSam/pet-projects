package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	start := `
_______________
File reader
_______________
Введите цифру для продолжения:
0. Exit
1. Чтение файла и вывод данных в консоль
2. Запись строки в файл
`

	read_file := `
Выберите способо чтения:
0. <Назад
1. Построчно(пакет bufioReader)
2. Целиком(пакет io)
`
	var command, level int
	for {
		fmt.Println(start)
		fmt.Scanln(&level)
		switch level {
		case 0:
			return
		case 1:
			fmt.Println(read_file)
			fmt.Scanln(&command)
			switch command {
			case 0:
				continue
			case 1:
				read_string()
				fmt.Scanln(&command)
			}
		case 2:
			write_file()
			fmt.Scanln(&command)
		}
	}
}

func read_string() {
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

func write_file() {
	var text string
	fmt.Println("Напишите текст для записи:")
	fmt.Scanln(&text)
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Данные записаны в файл file.txt")
}
