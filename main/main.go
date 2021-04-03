package main

import (
	r "filereader/utils/readfile"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", r.ReadAll("file.txt"))
}
