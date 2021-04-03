package utils

import (
	"io/ioutil"
	"log"
)

func ReadAll(f string) []byte {
	c, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
	return c
}

func ReadLines(c []byte) []string {
	line := ""
	lines := make([]string, 0)
	for _, char := range c {
		if string(char) == "\n" {
			log.Printf("%s\n", line)
			lines = append(lines, line)
			line = ""
		} else {
			line += string(c)
		}
	}
	return lines
}
