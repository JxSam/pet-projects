// Умный фильтр 😈
// Реализуйте фильтр с использованием анонимных функций, который способен выполнять следующие операции:

// Фильтровать массив целых чисел по заданным критериям:

// Сохранять только чётные числа (команда "even").
// Сохранять только нечётные числа ("odd").
// Сохранять только положительные числа ("positive").
// Сохранять только отрицательные числа ("negative").
// Сохранять только числа, которые являются квадратами целых чисел (например, 4, 9, 16, 25 и т.д.) ("squares").
// Сохранять только простые числа ("prime").
// Сохранять только числа, которые кратны 10 ("factor10").

// Команды фильтрации задаются в виде строковых ключей ("even", "odd", "positive", "negative", "squares", "prime", "factor10") в мапе filters. Вам необходимо дополнить эту мапу анонимными функциями, которые принимают число типа int и возвращают true или false в зависимости от логики выбранного фильтра.

// Формат входных данных:
// Программа должна принимать на вход три значения:

// Первое число n — количество чисел в массиве.
// Следующие n чисел — элементы массива.
// Одну строку operation — название команды фильтрации.
// Формат выходных данных:
// Программа должна вывести отфильтрованные числа массива, каждое через пробел. Если массив пуст после фильтрации, ничего выводить не нужно.

// Тестовые данные

// № Теста
// Входные данные
// Выходные данные
// 1

// 7
// 1 2 3 4 5 6 10
// even

// 2 4 6 10
// 2

// 5
// -3 -1 0 1 3
// positive

// 1 3
// 3

// 6
// 4 9 15 16 20 25
// squares

// 4 9 16 25
package function

import (
	"fmt"
	"math"
)

const (
	Even     = "even"
	Odd      = "odd"
	Positive = "positive"
	Negative = "negative"
	Squares  = "squares"
	Prime    = "prime"
	Factor10 = "factor10"
)

func filter(slice []int, filterFunc func(int) bool) []int {
	// менять эту функцию не нужно!!!
	// но я советую вам внимательно изучить эту функцию для саморазвития!

	// эта функция проходится по каждому элементу слайса
	// и отфильтровывает те, для которых filterFunc возвращает false

	result := make([]int, 0, len(slice))

	for _, num := range slice {
		if filterFunc(num) {
			result = append(result, num)
		}
	}
	return result
}

func MainFilter() {
	// дополните эту мапу своими анонимными функциями!
	filters := map[string]func(int) bool{
		Even: func(number int) bool {
			return number%2 == 0
		}, // Только чётные числа
		Odd: func(number int) bool {
			return number%2 != 0
		}, // Только нечётные числа
		Positive: func(number int) bool {
			return number > 0
		}, // Только положительные числа
		Negative: func(number int) bool {
			return number < 0
		}, // Только отрицательные числа
		Squares: func(n int) bool {
			if n < 0 {
				return false
			}
			root := int(math.Sqrt(float64(n)))
			return root*root == n
		}, // Только числа, которые являются квадратами
		Prime: func(number int) bool {
			if number <= 1 {
				return false
			}
			if number <= 3 {
				return true
			}
			if number%2 == 0 || number%3 == 0 {
				return false
			}

			limit := int(math.Sqrt(float64(number)))
			for i := 5; i <= limit; i += 6 {
				if number%i == 0 || number%(i+2) == 0 {
					return false
				}
			}
			return true
		}, // Только простые числа
		Factor10: func(number int) bool {
			return number%10 == 0
		}, // Только числа, которые кратны 10
	}

	// Ниже ничего менять не нужно
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	var operation string
	fmt.Scan(&operation)

	if filterFunc, exists := filters[operation]; exists {
		result := filter(numbers, filterFunc)
		for _, num := range result {
			fmt.Print(num, " ")
		}
	}
}
