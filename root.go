package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func get_float_input(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	var value float64
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Убираем лишние пробелы и переносы строк
		var err error
		value, err = strconv.ParseFloat(input, 64) // Пробуем преобразовать строку в число
		if err != nil {
			fmt.Println("Некорректный ввод, попробуйте еще раз.")
		} else {
			break
		}
	}
	return value
}

func find_roots(a float64, b float64, c float64) []float64 {
	D := b*b - 4.0*a*c
	roots := make([]float64, 0)
	if D == 0 {
		root := (-b) / (2.0 * a)
		roots = append(roots, root)
	} else if D > 0 {
		root1 := (-b + math.Sqrt(D)) / (2 * a)
		root2 := (-b - math.Sqrt(D)) / (2 * a)
		roots = append(roots, root1, root2)
	}
	return roots
}

func main() {
	a := get_float_input("Введите коэффициент A: ")
	b := get_float_input("Введите коэффициент B: ")
	c := get_float_input("Введите коэффициент C: ")

	roots := find_roots(a, b, c)
	if len(roots) == 0 {
		fmt.Println("Нет действительных корней.")
	} else {
		fmt.Println("Корни уравнения:", roots)
	}
}
