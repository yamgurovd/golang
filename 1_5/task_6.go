// Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.
package main

import "fmt"

func main() {
	var digit, hour, min_ uint

	fmt.Scan(&digit)

	hour = digit / 30
	min_ = 2 * (digit % 30)
	fmt.Println("It is", hour, "hours", min_, "minutes.")
}
