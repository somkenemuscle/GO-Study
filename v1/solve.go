package main

import "fmt"

// 🧪 Project 1 — “Smart Calculator CLI”

func solve(num1, num2 int, operator string) int {
  if num2 == 0 {
    fmt.Println("Cannot divide by zero")
    return 0
  }
	if operator == "+" {
		return num1 + num2
	} else if operator == "-" {
		return num1 - num2
	} else if operator == "*" {
		return num1 * num2
	} else if operator == "/" {
		return num1 / num2
	}else if operator == "%" {
		return num1 % num2
	}
  fmt.Println("Invalid operator")
	return 0
}

func main() {
	fmt.Println(solve(1, 0, "/"))
}
