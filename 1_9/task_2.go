/*По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	// Извлекаем цифры числа
	dig1 := num % 10         // Последняя цифра
	dig2 := (num / 10) % 10  // Средняя цифра
	dig3 := (num / 100) % 10 // Первая цифра

	// Проверяем, что все цифры различны
	if (dig1 != dig2) && (dig1 != dig3) && (dig2 != dig3) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
