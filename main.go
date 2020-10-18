package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	INPUT_LIST = "./input.yaml"
	SRC_DIR    = "./src"
	DST_DIR    = "./dist"
)

func main() {
	fileList := make([]string, 0)
	raw, err := ioutil.ReadFile(INPUT_LIST)
	if err != nil {
		exit(err)
	}
	err = yaml.Unmarshal(raw, &fileList)
	if err != nil {
		exit(err)
	}

	err = os.MkdirAll(DST_DIR, 0700)
	if err != nil {
		exit(err)
	}

	for _, file := range fileList {
		err := processFile(file)
		if err != nil {
			exit(err)
		}
	}
	os.Exit(0)
}
func exit(err error) {
	log.Printf("%+v\n", err)
	os.Exit(1)
}
