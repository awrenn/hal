package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

const (
	INPUT_LIST = "./input.yaml"
	SRC_DIR    = "./src"
	DST_DIR    = "./dist"
)

func main() {
	app := &cli.App{
		Name:   "hal",
		Usage:  "hal - compile html chunks into full files",
		Action: runHal,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}

func runHal(_c *cli.Context) error {
	fileList := make([]string, 0)
	raw, err := ioutil.ReadFile(INPUT_LIST)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, &fileList)
	if err != nil {
		return err
	}

	err = os.MkdirAll(DST_DIR, 0700)
	if err != nil {
		return err
	}

	for _, file := range fileList {
		err := processFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}
