package main

import (
	u "filereader/utils"
	"fmt"
)

func main() {
	bdepts := u.ReadAll("data/dept.data")
	ldepts := u.ReadLines(bdepts)
	depMap := u.GetUniqueDepartments(ldepts)
	fmt.Println(depMap)
}
