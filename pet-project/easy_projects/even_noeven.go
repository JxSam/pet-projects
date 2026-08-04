package easyprojects

import "fmt"

// Напишите функцию isEven, которая принимает одно целое число (int) и возвращает логическое значение (bool): true, если число является чётным, и false, если оно нечётное.

// Формат входных данных:
// Функция принимает один аргумент — целое число.

// Формат выходных данных:
// Функция возвращает логическое значение (bool).

func isEven(number int) bool {
	return number%2 == 0
}

func Even_noeven() {
	var input int
	fmt.Scan(&input)
	if isEven(input) {
		fmt.Println("Число четное!")
	} else {
		fmt.Println("Число нечетное!")
	}
}
