package config

var C config

type config struct {
	Switch struct {
		LastTab      bool   `yaml:"last_tab"`
		LastTabValue string `yaml:"last_tab_value"`
	}
}
