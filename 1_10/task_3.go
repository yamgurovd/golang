// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
// Программа в первой строке получает на вход число n - количество чисел в последовательности,
// во второй строке -- n чисел, входящих в данную последовательность.
package main

import "fmt"

func main() {
	var n, digit, sum, i int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&digit)
		sum += digit
	}
	fmt.Println(sum)
}
