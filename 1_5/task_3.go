// По данному целому числу, найдите его квадрат.
//
//Формат входных данных
//На вход дается одно целое число.
//
//Формат выходных данных
//Программа должна вывести квадрат данного числа.

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	b = a * a
	fmt.Println(b)
}
