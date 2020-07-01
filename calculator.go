package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entry string, operator string) {
	entry_split := strings.Split(entry, " ")
	operator1 := cast(entry_split[0])
	operator2 := cast(entry_split[1])
	switch operator {
	case "+":
		fmt.Println(operator1 + operator2)
	case "-":
		fmt.Println(operator1 - operator2)
	case "*":
		fmt.Println(operator1 * operator2)
	case "/":
		fmt.Println(operator1 / operator2)
	default:
		fmt.Println(operator, "Operator not supported")
	}
}

func cast(operator string) int {
	operator_cast, _ := strconv.Atoi(operator)
	return operator_cast
}

func main(){
	fmt.Println("Enter the numbers to operate with a space between (1 2)")
	entry := read_entry()
	fmt.Println("Enter the operation (+ - * /)")
	operator := read_entry()
	c := calc{}
	c.operate(entry, operator)
}

func read_entry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}