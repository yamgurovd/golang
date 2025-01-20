// Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
//
// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.
//
// Формат выходных данных
// Выведите одно целое число - первую цифру заданного числа.

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case n < 10:
		fmt.Println(n)
	case n < 100:
		fmt.Println(n / 10)
	case n < 1000:
		fmt.Println(n / 100)
	case n < 10000:
		fmt.Println(n / 1000)
	case n < 100000:
		fmt.Println(n / 10000)
	}
}
