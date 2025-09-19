package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateAverage(input string) string {
	fields := strings.Fields(input)
	var sum int
	var count int
	for _, f := range fields {
		// Проверяем, целое ли число
		n, err := strconv.Atoi(f)
		if err != nil {
			continue
		}
		if n >= 0 && n <= 100 {
			sum += n
			count++
		}
	}
	if count == 0 {
		return "Нет корректных оценок"
	}
	avg := float64(sum) / float64(count)
	// Округление до 1 знака после запятой математически
	avg = math.Round(avg*10) / 10
	return fmt.Sprintf("%.1f", avg)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	res := calculateAverage(input)

	fmt.Println(res)
}
