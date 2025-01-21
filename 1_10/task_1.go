// Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится
// в новой строке.
package main

import "fmt"

func main() {
	var digit, square int
	for digit := 1; digit < 11; digit++ {
		square = digit * digit
		fmt.Print(square)
	}
}
