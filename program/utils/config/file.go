package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/bsthun/switchbox/program/entries/val"
	"github.com/bsthun/switchbox/program/wrappers/interactive"
)

func loadConfigFile() (cfg *config) {
	// * Read YAML configuration file
	var yml []byte
	if y, err := ioutil.ReadFile("config.yml"); err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO READ YAML CONFIGURATION FILE", err, nil)
	} else {
		yml = y
	}

	if err := yaml.Unmarshal(yml, cfg); err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO PARSE YAML CONFIGURATION FILE", err, nil)
	}

	return
}
