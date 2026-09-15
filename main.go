package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("请输入第一个数字: ")
	fmt.Scanln(&num1)

	fmt.Print("请输入运算符 (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("请输入第二个数字: ")
	fmt.Scanln(&num2)

	// 其他逻辑
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("错误: 除数不能为零")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		}
	default:
		fmt.Println("错误: 无效的运算符")
	}
}