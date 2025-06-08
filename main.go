package main

import (
	"fmt"
	"strings"
)

func Rating_stars(score int) string {
	if score < 0 || score > 100 {
		return "☆☆☆☆☆"
	}

	stars := score / 20
	return strings.Repeat("★", stars) + strings.Repeat("☆", 5-stars)
}

func main() {
	var score int

	fmt.Println("Введіть число від 0 до 100:")
	fmt.Scanln(&score)

	result := Rating_stars(score)
	fmt.Println("Рейтинг:", result)

}
