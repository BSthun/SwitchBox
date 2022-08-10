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
	if _yml, _err := ioutil.ReadFile("config.yml"); _err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO READ YAML CONFIGURATION FILE", _err)
	} else {
		yml = _yml
	}

	if _err := yaml.Unmarshal(yml, cfg); _err != nil {
		interactive.Throw(val.DefaultTag, "UNABLE TO PARSE YAML CONFIGURATION FILE", _err)
	}

	return
}
