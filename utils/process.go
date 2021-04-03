package utils

import (
	"strconv"
	"strings"
)

func GetUniqueDepartments(d []string) map[int]string {
	depMap := make(map[int]string, 0)
	for _, line := range d {
		cIndex := strings.Index(line, ",")
		num, _ := strconv.Atoi(line[:cIndex])
		depMap[num] = line[cIndex+1:]
	}
	return depMap
}
