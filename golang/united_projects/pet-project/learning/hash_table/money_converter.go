package hashtable

import (
	"fmt"
	"strconv"
)

func Convert() {
	var value string
	var number, action int
	// Курс валют относительно 1 USD
	var rates = map[string]float64{
		"USD": 1.0,   // Доллар США
		"EUR": 0.92,  // Евро
		"RUB": 90.0,  // Российский рубль
		"JPY": 157.0, // Японская иена
		"CNY": 7.25,  // Китайский юань
		"GBP": 0.78,  // Британский фунт
		"KZT": 460.0, // Казахстанский тенге
		"TRY": 32.5,  // Турецкая лира
		"INR": 83.0,  // Индийская рупия
		"BRL": 5.12,  // Бразильский реал
		"AUD": 1.50,  // Австралийский доллар
		"CAD": 1.36,  // Канадский доллар
		"CHF": 0.89,  // Швейцарский франк
		"SEK": 10.8,  // Шведская крона
		"NOK": 10.5,  // Норвежская крона
	}
	var currency_map = map[int]string{
		1:  "USD", // Доллар США
		2:  "EUR", // Евро
		3:  "RUB", // Российский рубль
		4:  "JPY", // Японская иена
		5:  "CNY", // Китайский юань
		6:  "GBP", // Британский фунт
		7:  "KZT", // Казахстанский тенге
		8:  "TRY", // Турецкая лира
		9:  "INR", // Индийская рупия
		10: "BRL", // Бразильский реал
		11: "AUD", // Австралийский доллар
		12: "CAD", // Канадский доллар
		13: "CHF", // Швейцарский франк
		14: "SEK", // Шведская крона
		15: "NOK", // Норвежская крона
	}
	currency_text := "Доступные валюты для конвертации:\n" +
		"1. NOK\n" +
		"2. RUB\n" +
		"3. KZT\n" +
		"4. AUD\n" +
		"5. CHF\n" +
		"6. SEK\n" +
		"7. INR\n" +
		"8. BRL\n" +
		"9. GBP\n" +
		"10. CAD\n" +
		"11. USD\n" +
		"12. JPY\n" +
		"13. TRY\n" +
		"14. EUR\n" +
		"15. CNY"
	fmt.Println("Добро пожаловать! Это конвертер валют\n" + currency_text)
	for {
		fmt.Println(currency_text)
		fmt.Println("Введите сумму в USD:")
		fmt.Scan(&value)
		value_correct, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("Неверное число. Повторите попытку")
		} else {
			fmt.Println("Выберите номер валюты для конвертации из списка выше:")
			fmt.Scan(&number)
			fmt.Printf("%s USD = %f %s\n", value, value_correct*rates[currency_map[number]], currency_map[number])
		}
		fmt.Println("Выберите действие. Выход - 1, Далее - enter")
		fmt.Scanln(&action)
		if action == 1 {
			break
		} else {
			continue
		}
	}
}
