package mycalculator

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

func read_entry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}