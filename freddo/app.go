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

type Branch struct {
	Script      string
	UpdateQueue chan string
}

type App struct {
	Name     string
	Branches map[string][]*Branch
	Secret   []byte
	Script   string
}

func NewBranch(script string) *Branch {
	return &Branch{
		Script:      script,
		UpdateQueue: make(chan string, 1),
	}
}

func NewApp(name string) *App {
	app := &App{
		Branches: make(map[string][]*Branch),
		Name:     name,
	}
	return app
}

func (a *App) AddBranch(ref, script string) *Branch {
	branch := NewBranch(script)
	_, ok := a.Branches[ref]
	if !ok {
		a.Branches[ref] = []*Branch{}
	}
	a.Branches[ref] = append(a.Branches[ref], branch)
	return branch
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

func (b *Branch) QueueUpdate() error {
	select {
	case b.UpdateQueue <- "update":
		return nil
	default:
		return errors.New("Could not queue update for " + b.Script)
	}
}

func (b *Branch) LoopQueue() {
	for _ = range b.UpdateQueue {
		b.Update()
	}
}

func (b *Branch) Update() {
	log.Print("Running: ", b.Script)

	out, err := b.RunScript()
	if err != nil {
		log.Print("Failed: ", b.Script)
		log.Print("output: ", string(out))
		log.Print(err)
		return
	}
	log.Print("Finished: ", b.Script)
}

func (b *Branch) RunScript() ([]byte, error) {
	c := exec.Command("/bin/sh", "-c", b.Script)
	return c.CombinedOutput()
}
