package main

import (
	"fmt"
)

var elem1, operand, elem2 string
var isRoman bool = false

// словарь перевода римских чисел в int
var romanDict = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"VIV":  9,
	"X":    10,
}

// словарь перевода арабских чисел в int
var arabDict = map[string]int{
	"1":    1,
	"2":    2,
	"3":    3,
	"4":    4,
	"5":    5,
	"6":    6,
	"7":    7,
	"8":    8,
	"9":    9,
	"10":   10,
}

// проверка корректности ввода чисел
func isCorrect(elem1 string, elem2 string, operand string) (int, int, string){
	x_arab, ok_x_arab := arabDict[elem1]
	y_arab, ok_y_arab := arabDict[elem2]
	
	x_roman, ok_x_roman := romanDict[elem1]
	y_roman, ok_y_roman := romanDict[elem2]
	
	if ok_x_arab && ok_y_arab {
        return x_arab, y_arab, operand
	} else if ok_x_roman && ok_y_roman {
	    isRoman = true
	    return x_roman, y_roman, operand
	} else {
	    panic("Некорректный ввод чисел")
	}
}

// ф-ция сложения
func add(x int, y int) (int) {
    return x + y
}

// ф-ция вычитания
func subtraction(x int, y int) (int) {
    if isRoman {
        if x - y <= 0 {
            panic("Римские числа не могут быть отрицательными и 0")
        } else {
            return x - y
        }
    } else {
        return x - y
    }
}

// ф-ция умножения
func multiply(x int, y int) (int) {
    return x * y
}

// ф-ция деления
func division(x int, y int) (int) {
    if y == 0 {
        panic("Нельзя делить на ноль")
    } else {
        return int(x / y)
    }
}

// ф-ция вычисления + проверка корректности ввода операнда
func calc(x int, y int, operand string) (int) {
    var result int
    switch operand {
    case "+": result = add(x, y)
    case "-": result = subtraction(x, y)
    case "*": result = multiply(x, y)
    case "/": result = division(x, y)
    default: fmt.Println("Неизвестный операнд")
    }
    return result
}

// ф-ция вывода результат вычисления
func toPrint(isRoman bool, elem1 string, elem2 string, operand string, result int) {
    if isRoman {
        text := toRomanPrint(result)
        fmt.Println(elem1, operand, elem2, "=", text)
    } else {
        fmt.Println(elem1, operand, elem2, "=", result)
    }
}

// ф-ция перевода результата вычисления в римские числа
func toRomanPrint(number int) (string) {
    var romanToArabDict = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
	10: "X",
	11: "XI",
	12: "XII",
	13: "XIII",
	14: "XIV",
	15: "XV",
	16: "XVI",
	17: "XVII",
	18: "XVIII",
	19: "XIX",
	20: "XX",
    }
    text, _ := romanToArabDict[number]
    return text
}

func main() {
    
    fmt.Println(`
        Калькулятор для арабских и римских чисел.
        Следует вводить числа от 1 до 10 и от I до X
        в формате x + y.
    `)
	
	startFlag := true
	for startFlag {
	    // ввод и считываение выражения
	    fmt.Println("Введите выражение:")
    	_, _ = fmt.Scan(&elem1, &operand, &elem2)
    	
    	// проверка корректности ввода чисел
    	x, y, operand := isCorrect(elem1, elem2, operand)
    	
    	// вычисление результата
    	result := calc(x, y, operand)
    	
    	// вывод результата
    	toPrint(isRoman, elem1, elem2, operand, result)
    	
    	// повторный запрос на вычисление
    	fmt.Println("Начать заново? да/нет")
    	var nextStep string
    	_,_ = fmt.Scan(&nextStep)
    	if nextStep == "нет" {
    	    startFlag = false
    	}
	}
	
}

