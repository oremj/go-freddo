package freddo

import (
	"log"
	"os/exec"
	"sync"

	"github.com/oremj/go-freddo/semaphore"
)

type App struct {
	updateLock sync.Mutex
	waitSemaphore *semaphore.Semaphore
	Name       string
	Script     string
}

func NewApp(name string) *App {
	app := &App{
		Name:    name,
		waitSemaphore: semaphore.NewSemaphore(1),
	}
	return app
}

func (a *App) Update() {
	if !a.waitSemaphore.Wait() {
		log.Println("There are already updates waiting.")
		return
	}

	a.updateLock.Lock()
	defer a.updateLock.Unlock()
	a.waitSemaphore.Signal()

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
