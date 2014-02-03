package freddo

import (
	"encoding/json"
	"os"
)

type Freddo struct {
	Apps map[string]*App
}

type ConfigJson struct {
	Apps map[string]struct {
		Script string `json:"script"`
	} `json:"apps"`
}

// Read config and return freddo object.
func NewFreddo(configFile string) (*Freddo, error) {
	f, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	c := new(ConfigJson)
	err = decoder.Decode(c)
	if err != nil {
		return nil, err
	}

	freddo := &Freddo{
		Apps: make(map[string]*App),
	}

	for app, val := range c.Apps {
		tmp := NewApp()
		tmp.Script = val.Script
		freddo.Apps[app] = tmp
	}
	return freddo, nil
}
