package main

import (
	r "filereader/utils"
)

func main() {
	con := r.ReadAll("data/emp.data")
	r.GetDepartment(con)
}
