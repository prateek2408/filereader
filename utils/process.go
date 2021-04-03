package utils

import (
	"strconv"
	"strings"
)

type Edata struct {
	ID       string
	Name     string
	Salary   string
	Dept     int
	DeptName string
}

func GetUniqueDepartments(d []string) map[int]string {
	depMap := make(map[int]string, 0)
	for _, line := range d {
		cIndex := strings.Index(line, ",")
		num, _ := strconv.Atoi(line[:cIndex])
		depMap[num] = line[cIndex+1:]
	}
	return depMap
}

func fillStruct(d []byte) Edata {
	var data Edata
	raw_data = strings.Split(string(d), ",")
	data.
	return data
}

func GenerateEmployeeMemoryData(d []byte) []Edata {
	var temp []byte
	var data []Edata
	for _, bs := range d {
		if string(bs) == "\n" {
			dt := fillStruct(bs)
			data := append(data, dt)
			temp = 0
		} else {
			temp += bs
		}
	}
	return data

}
