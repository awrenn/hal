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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "src",
				Value: SRC_DIR,
				Usage: "Src directory to find html snippets in",
			},
			&cli.StringFlag{
				Name:  "dst",
				Value: DST_DIR,
				Usage: "Dst directory to write output html to",
			},
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   INPUT_LIST,
				Usage:   "Input yaml file to read the input from",
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}

func runHal(c *cli.Context) error {
	inputConf := InputConfig{}
	raw, err := ioutil.ReadFile(c.String("input"))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, &inputConf)
	if err != nil {
		return err
	}

	conf := HalConfig{
		Dst: c.String("dst"),
		Src: c.String("src"),
	}
	err = os.MkdirAll(conf.Dst, 0700)
	if err != nil {
		return err
	}

	for _, file := range inputConf.Targets {
		err := processFile(conf, file.Src, file.Dst)
		if err != nil {
			return err
		}
	}
	return nil
}
