package config

import (
	"encoding/json"
	"io/ioutil"
)

type cfg struct {
	Processs []*Process `json:"processes"`
}

type Process struct {
	Name     string `json:"name"`
	RCmd     string `json:"rcmd"`
	Endpoint string `json:"endpoint"`

	CheckInterval  int64 `json:"checkInterval"`
	ReloadInterval int64 `json:"reloadInterval"`
}

func Load(file string) (process []*Process, err error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	var c cfg

	err = json.Unmarshal(dat, &c)
	if err != nil {
		return
	}

	process = c.Processs
	return
}
