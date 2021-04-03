package main

import (
	u "filereader/utils"
	"fmt"
)

func main() {
	//Fetching data from departments.
	bdepts := u.ReadAll("data/dept.data")
	ldepts := u.ReadLines(bdepts)
	depMap := u.GetUniqueDepartments(ldepts)
	//Fetching data from Emp File.
	edepts := u.ReadAll("data/emp.data")
	empStruct := u.GenerateEmployeeMemoryData(edepts)

}
