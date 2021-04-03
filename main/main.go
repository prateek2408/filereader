package main

import (
	r "filereader/utils"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", r.ReadAll("data/emp.data"))
}
