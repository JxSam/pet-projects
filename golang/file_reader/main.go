package main

import (
	"bufio"
	"fmt"
	"io"
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
3. Добавить текст в файл
4. Узнать размер файла
5. Объединить 3 файла в один
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
				level = 0
			case 1:
				read_string()
				fmt.Scanln(&command)
			case 2:
				read_files()
				fmt.Scanln(&command)
			}
		case 2:
			write_file()
			fmt.Scanln(&command)
		case 3:
			add_file()
			fmt.Println("Нажмите enter для продолжения")
			fmt.Scanln(&command)
		case 4:
			sizing_file()
			fmt.Scanln(&command)
		case 5:
			merge_files()
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

func read_files() {
	var path string

	fmt.Println("Укажите путь до файла:")
	fmt.Scanln(&path)

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}

	defer file.Close()

	text, err := io.ReadAll(file)

	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println(string(text))
}

func write_file() {
	var text, nameFile string
	fmt.Println("Напишите название файла(без расширения файла):")
	fmt.Scanln(&nameFile)
	fmt.Println("Напишите текст для записи:")
	fmt.Scanln(&text)
	file, err := os.Create(nameFile + ".txt")
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

func add_file() {
	var text, path string

	fmt.Println("Укажите путь до файла:")
	fmt.Scanln(&path)

	fmt.Println("Напишите текст:")
	fmt.Scanln(&text)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(text)

	fmt.Println("Текст успешно добавлен в файл!")
}

func sizing_file() {
	var path string

	fmt.Println("Укажите путь до файла:")
	fmt.Scanln(&path)

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}

	defer file.Close()

	text, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(text), "bytes")
	fmt.Println(float64(len(text))/1024, "kilobytes")
	fmt.Println(float64(len(text))/1048576, "megabytes")
}

func merge_files() {
	var path1, path2, path3 string

	fmt.Println("Файл 1:")
	fmt.Scanln(&path1)
	file, err := os.Open(path1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("Файл 2:")
	fmt.Scanln(&path2)
	file2, err := os.Open(path2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	fmt.Println("Файл 3:")
	fmt.Scanln(&path3)
	file3, err := os.Open(path3)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file3.Close()

	output, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer output.Close()

	text, _ := io.ReadAll(file)
	text2, _ := io.ReadAll(file2)
	text3, _ := io.ReadAll(file3)

	output.WriteString(string(text) + string(text2) + string(text3))

	fmt.Println("Файлы соединены в output.txt")
}
