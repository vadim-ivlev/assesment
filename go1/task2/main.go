package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AnalyzeGrades(input string) []string {
	fields := strings.Fields(input)
	counts := make([]int, 5)
	total := 0
	for _, f := range fields {
		var n int
		_, err := fmt.Sscanf(f, "%d", &n)
		if err != nil || n <= 0 || n > 100 {
			continue
		}
		switch {
		case n <= 20:
			counts[0]++
		case n <= 40:
			counts[1]++
		case n <= 60:
			counts[2]++
		case n <= 80:
			counts[3]++
		case n <= 100:
			counts[4]++
		}
		total++
	}
	if total == 0 {
		return []string{"Нет корректных оценок"}
	}
	ranges := [5]string{"0–20", "21–40", "41–60", "61–80", "81–100"}
	res := make([]string, 5)
	for i := 0; i < 5; i++ {
		percent := float64(counts[i]) * 100 / float64(total)
		res[i] = fmt.Sprintf("%s: %d оценок, %.1f%%", ranges[i], counts[i], percent)
	}

	// res := []string{
	// 	fmt.Sprintf("0–20: %d оценок, %.1f%%", counts[0], float64(counts[0])*100/float64(total)),
	// 	fmt.Sprintf("21–40: %d оценок, %.1f%%", counts[1], float64(counts[1])*100/float64(total)),
	// 	fmt.Sprintf("41–60: %d оценок, %.1f%%", counts[2], float64(counts[2])*100/float64(total)),
	// 	fmt.Sprintf("61–80: %d оценок, %.1f%%", counts[3], float64(counts[3])*100/float64(total)),
	// 	fmt.Sprintf("81–100: %d оценок, %.1f%%", counts[4], float64(counts[4])*100/float64(total)),
	// }

	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	for _, line := range AnalyzeGrades(input) {
		fmt.Println(line)
	}
}
