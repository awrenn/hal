package main

type HalConfig struct {
	Src string
	Dst string
}

type InputConfig struct {
	Targets []InputItem `yaml:"targets"`
}

type InputItem struct {
	Src string
	Dst string
}
