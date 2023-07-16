package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (int, error) {
	var sum int
	var greatest int
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		num := romanNumerals[rune(letter)]
		if num >= greatest {
			greatest = num
			sum = sum + num
			continue
		}
		sum = sum - num
	}
	return sum, nil
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func getAndCheckNums(c []string) (int, int, error, bool) {
	var check bool
	var check1 string
	var check2 string
	var err error
	var num1int int
	var num2int int

	if c[0] == strings.ToLower(c[0]) {
		check1 = "arab"
		num1int, _ = strconv.Atoi(c[0])
	} else {
		check1 = "rom"
		num1int, err = romanToInt(c[0])
		check = true
	}
	if c[2] == strings.ToLower(c[2]) {
		check2 = "arab"
		num2int, _ = strconv.Atoi(c[2])
	} else {
		check2 = "rom"
		num2int, err = romanToInt(c[2])
		check = true
	}
	if check1 != check2 {
		return 0, 0, errors.New("Ошибка: Используются одновременно разные системы счисления."), false
	}
	if num1int > 10 || num2int > 10 {
		return 0, 0, errors.New("Ошибка: А цыфорки ты большие выбрал, скромнее надо быть!"), false
	}
	return num1int, num2int, err, check
}

func processStack(s []string) (int, error, bool) {
	var result int
	fmt.Println(len(s))
	if len(s)-1 < 2 {
		return 0, errors.New("Ошибка: Cтрока не является математической операцией."), false
	} else if len(s) > 3 {
		return 0, errors.New("Ошибка: Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."), false
	}
	num1, num2, err, check := getAndCheckNums(s)
	if err != nil {
		return 0, err, false
	}
	switch s[1] {
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
		if check && result <= 0 {
			return 0, errors.New("Ошибка: В римской системе нет отрицательных чисел."), false
		}
	}

	return result, nil, check
}

func main() {
	var expressions []string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("gocalc>")
		for scanner.Scan() {
			expressions = strings.Split(scanner.Text(), " ")
			res, err, check := processStack(expressions)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			} else {
				if check {
					fmt.Println(intToRoman(res))
				} else {
					fmt.Println(res)
				}
			}
			fmt.Print("gocalc>")
		}
	}
}
