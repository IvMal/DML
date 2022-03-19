package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите сумму дохода")
	var income, tax float64
	fmt.Scan(&income)
	switch {
	case income >= 70000:
		tax = (income - 70000) * 50 / 100
		fmt.Println("Сумма налога свыше 70000 составляет 50% и равен:", tax)
		income = 70000
		fallthrough
	case income >= 50000:
		tax = (income - 50000) * 30 / 100
		fmt.Println("Сумма налога свыше 50000 и меньше 70000 составляет 30% и равен:", tax)
		income = 50000
		fallthrough
	case income >= 10000:
		tax = (income - 10000) * 20 / 100
		fmt.Println("Сумма налога свыше 10000 и до 50000 составляет 20% и равен:", tax)
		income = 10000
		fallthrough
	case income > 0:
		tax = income * 13 / 100
		fmt.Println("Сумма налога до 10000 13% и равен:", tax)
	default:
		fmt.Println("Доход должен быть положительным")
	}
}
