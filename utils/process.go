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

func fillStruct(d []byte, dm map[int]string) Edata {
	var data Edata
	raw_data := strings.Split(string(d), ",")
	data.ID = raw_data[0]
	data.Name = raw_data[1]
	data.Salary = raw_data[2]
	di, _ := strconv.Atoi(raw_data[3])
	data.Dept = di
	data.DeptName = dm[di]
	return data
}

func GenerateEmployeeMemoryData(d []byte, dm map[int]string) []Edata {
	var temp []byte
	var data []Edata
	for _, bs := range d {
		if string(bs) == "\n" {
			dt := fillStruct(temp, dm)
			data = append(data, dt)
			temp = temp[:0]
		} else {
			temp = append(temp, bs)
		}
	}
	return data

}
