package freddo

import "github.com/BurntSushi/toml"

type Freddo struct {
	Apps map[string]*App
}

type ConfigTOML struct {
	Apps map[string]struct {
		Secret string
		Branch []struct {
			Ref    string
			Script string
		}
	}
}

// Read config and return freddo object.
func NewFreddo(configFile string) (*Freddo, error) {
	c := new(ConfigTOML)
	_, err := toml.DecodeFile(configFile, c)
	if err != nil {
		return nil, err
	}

	freddo := &Freddo{
		Apps: make(map[string]*App),
	}

	for app, val := range c.Apps {
		tmp := NewApp(app)
		tmp.Secret = []byte(val.Secret)
		for _, branch := range val.Branch {
			b := tmp.AddBranch(branch.Ref, branch.Script)
			go b.LoopQueue()
		}
		freddo.Apps[app] = tmp
	}
	return freddo, nil
}
