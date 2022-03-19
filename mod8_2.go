/*Пользователь вводит будний день недели в сокращённой
форме (пн, вт, ср, чт, пт) и получает развёрнутый список
всех последующих рабочих дней, включая пятницу.*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите будний день недели в сокращённом виде:")
	var dayShortText string
	fmt.Scan(&dayShortText)
	switch dayShortText{
	case "пн":
		fmt.Println("Понедельник")
		fallthrough
	case "вт":
		fmt.Println("Вторник")
		fallthrough
	case "ср":
		fmt.Println("Среда")
		fallthrough
	case "чт":
		fmt.Println("Четверг")
		fallthrough
	case "пт":
		fmt.Println("Пятница")
	default:
		fmt.Println("Попробуйте ещё раз")
	}
}
