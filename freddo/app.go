package freddo

import (
	"os/exec"
	"sync"
)

type App struct {
	*sync.Mutex
	Script string
}

func NewApp() *App {
	app := &App{
		Mutex: new(sync.Mutex),
	}
	return app
}

func (a *App) RunScript() ([]byte, error) {
	c := exec.Command("/bin/sh", "-c", a.Script)
	return c.CombinedOutput()
}
