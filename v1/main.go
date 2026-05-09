package main

import (
	"fmt"
	"strconv"
)

// 🧪 Project 4 — “Number Analyzer System”

func checkParity(num int) string {
	if num%2 == 0 {
		return strconv.Itoa(num) + " is an even number"
	}
	return strconv.Itoa(num) + " is an odd number"
}

func checkSign(num int) string {
	if num == 0 {
		return strconv.Itoa(num) + " is zero"
	}
	if num > 0 {
		return strconv.Itoa(num) + " is positive"
	}
	return strconv.Itoa(num) + " is negative"
}

func compareToHundred(num int) string {
	if num < 100 {
		return strconv.Itoa(num) + " is less than 100"
	}
	if num > 100 {
		return strconv.Itoa(num) + " is greater than 100"
	}
	return strconv.Itoa(num) + " is equal to 100"
}

func analyzeNumber(num int) string {
	parity := checkParity(num)
	integer := checkSign(num)
	evaluation := compareToHundred(num)
	return parity + "\n" + integer + "\n" + evaluation

}

func main() {
	fmt.Println(analyzeNumber(100))
	fmt.Println(analyzeNumber(95))
	fmt.Println(analyzeNumber(0))
	fmt.Println(analyzeNumber(-40))
}

// 🧪 Project 1 — “Smart Calculator CLI”

// func solve(num1, num2 int, operator string) int {
// 	if num2 == 0 {
// 		fmt.Println("Cannot divide by zero")
// 		return 0
// 	}
// 	if operator == "+" {
// 		return num1 + num2
// 	} else if operator == "-" {
// 		return num1 - num2
// 	} else if operator == "*" {
// 		return num1 * num2
// 	} else if operator == "/" {
// 		return num1 / num2
// 	} else if operator == "%" {
// 		return num1 % num2
// 	}
// 	fmt.Println("Invalid operator")
// 	return 0
// }

// 🧪 Project 2 — “Grade Calculator”

// func grade(score int) string {
// 	if score < 0 || score > 100 {
// 		return "Invalid score"
// 	}

// 	if score >= 90 {
// 		return "A"
// 	}

// 	if score >= 80 {
// 		return "B"
// 	}

// 	if score >= 70 {
// 		return "C"
// 	}

// 	if score >= 60 {
// 		return "D"
// 	}

// 	return "F"
// }

// 🧪 Project 3 — “Login Authentication”

// var uname = "admin"
// var pwd = "secret123"

// func login(username, password string) string {
// 	if username == "" || password == "" {
// 		return "Username and password required"
// 	}
// 	if username != uname && password != pwd {
// 		return "Invalid username and password"
// 	}
// 	if username != uname {
// 		return "Invalid username"
// 	}
// 	if password != pwd {
// 		return "Invalid password"
// 	}

// 	return "Login successful"

// }
