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
