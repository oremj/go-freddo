package freddo

import (
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"fmt"
	"hash"
	"log"
	"os/exec"
	"strings"
)

type App struct {
	Name        string
	Secret      []byte
	Script      string
	UpdateQueue chan string
}

func NewApp(name string) *App {
	app := &App{
		Name:        name,
		UpdateQueue: make(chan string, 1),
	}
	return app
}

func (a *App) HmacEqual(msg []byte, msgSig string) (bool, error) {
	parts := strings.Split(msgSig, "=")
	if len(parts) != 2 {
		return false, errors.New("Invalid signature: " + msgSig)
	}
	var mac hash.Hash
	switch parts[0] {
	case "sha1":
		mac = hmac.New(sha1.New, a.Secret)
	default:
		return false, errors.New("Unsupported hash type: " + parts[0])
	}
	mac.Write(msg)
	expected := mac.Sum(nil)
	return parts[1] == fmt.Sprintf("%x", expected), nil
}

func (a *App) QueueUpdate() error {
	select {
	case a.UpdateQueue <- "update":
		return nil
	default:
		return errors.New("Could not queue update for " + a.Name)
	}
}

func (a *App) LoopQueue() {
	for _ = range a.UpdateQueue {
		a.Update()
	}
}

func (a *App) Update() {
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
