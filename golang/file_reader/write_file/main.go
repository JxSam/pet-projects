package main

import (
	"fmt"
	"os"
)

const Text = `New document file
Test
test
test
Hello, World!
Hello, World!
Hello, World!
Hello, World!
Hello, World!
Hello, World!
Bye, bye!
`

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(Text)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Данные записаны в файл")
}
