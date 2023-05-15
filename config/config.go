package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var (
	gData []byte
	err   error
)

func Load(obj any) error {
	if len(gData) < 1 {
		gData, err = os.ReadFile("./application.yml")
		if err != nil {
			return err
		}
	}
	return yaml.Unmarshal(gData, obj)
}
