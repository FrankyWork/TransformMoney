package main

import "fmt"

func main() {
	var firstMoney string
	var lastMoney string
	var haveMoney float64
	cursD := [2]float64{0: 0.91, 1: 124}
	cursE := [2]float64{0: 135, 1: 1.09}
	cursR := [2]float64{0: 0.0081, 1: 0.0074}

	fmt.Println("Введите вашу валюту: (EUR) (DOL) (RUB):")
	fmt.Scan(&firstMoney)
	fmt.Println("Введите желаемую валюту: (EUR) (DOL) (RUB):")
	fmt.Scan(&lastMoney)
	fmt.Println("Введите сумму перевода")
	fmt.Scan(&haveMoney)

	if firstMoney == lastMoney {
		fmt.Println(haveMoney)
	} else if firstMoney == "EUR" && lastMoney == "RUB" {
		fmt.Println(haveMoney, "Евро", "=", (haveMoney * cursE[0]), "Руб.")
	} else if firstMoney == "EUR" && lastMoney == "DOL" {
		fmt.Println(haveMoney, "Евро", "=", (haveMoney * cursE[1]), "Дол.")
	} else if firstMoney == "RUB" && lastMoney == "EUR" {
		fmt.Println(haveMoney, "Руб.", "=", (haveMoney * cursR[1]), "Евр.")
	} else if firstMoney == "RUB" && lastMoney == "DOL" {
		fmt.Println(haveMoney, "Руб.", "=", (haveMoney * cursR[0]), "Дол.")
	} else if firstMoney == "DOL" && lastMoney == "EUR" {
		fmt.Println(haveMoney, "Долларов", "=", (haveMoney * cursD[0]), "Евр.")
	} else if firstMoney == "DOL" && lastMoney == "RUB" {
		fmt.Println(haveMoney, "Долларов", "=", (haveMoney * cursD[1]), "Руб.")
	}
}
