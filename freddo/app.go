package freddo

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

type App struct {
	updateLock *sync.Mutex
	waitLock   *sync.Mutex
	waiting    int
	Name       string
	Script     string
}

func NewApp(name string) *App {
	app := &App{
		Name:       name,
		updateLock: new(sync.Mutex),
		waitLock:   new(sync.Mutex),
	}
	return app
}

func (a *App) incWaitLock() error {
	a.waitLock.Lock()
	defer a.waitLock.Unlock()

	if a.waiting > 0 {
		return fmt.Errorf("There are already %d updates waiting.", a.waiting)
	}

	a.waiting++
	return nil
}

func (a *App) decWaitLock() {
	a.waitLock.Lock()
	defer a.waitLock.Unlock()

	a.waiting--
}

func (a *App) Update() {
	err := a.incWaitLock()
	if err != nil {
		log.Print(err)
		return
	}

	a.updateLock.Lock()
	defer a.updateLock.Unlock()
	a.decWaitLock()

	log.Print("Running: ", a.Script)

	out, err := a.RunScript()
	if err != nil {
		log.Print("Failed: ", a.Script)
		log.Print("output: ", string(out))
		log.Print(err)
	}
}

func (a *App) RunScript() ([]byte, error) {
	c := exec.Command("/bin/sh", "-c", a.Script)
	return c.CombinedOutput()
}
