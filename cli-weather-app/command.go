package main

import (
	"flag"
	"fmt"
)

type CmdFlags struct {
	Fetch string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Fetch, "fetch", "", "Fetch weather data for a city like fetch city_name")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute() {
	switch {
	case cf.Fetch != "":
		Fetch(cf.Fetch)
	default:
		fmt.Println("Invalid command")

	}
}
