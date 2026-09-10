package main

import "fmt"

const start = `
_______________
File reader
_______________
Введите цифру для продолжения:

1. Чтение файла и вывод данных в консоль
2. Запись строки в файл
`

const read_file = `
Выберите способо чтения:
1. Построчно(пакет bufioReader)
2. Целиком(пакет io)
`

func main() {
	var command int

	fmt.Println(start)
	fmt.Scanln(&command)

	if command == 1 {
		fmt.Println(read_file)
	}
}
