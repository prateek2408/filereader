package readfile

import (
	"io/ioutil"
	"log"
)

func ReadAll(f string) {
	c, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalln(err)
	}
}
