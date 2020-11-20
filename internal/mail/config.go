package mail

import (
	"encoding/json"
	"io/ioutil"
)

func LoadConfig(filepath string) (Config, error) {
	cfg := new(Config)
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(b, cfg)
	if err != nil {
		return Config{}, err
	}
	return *cfg, nil
}
