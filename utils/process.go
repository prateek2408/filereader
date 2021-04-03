package utils

import "fmt"

func GetDepartment(c []byte) {
	for _, line := range c {
		fmt.Print(string(line))
		if string(line) == "\n" {
			fmt.Println()
		}
	}
}
