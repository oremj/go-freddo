package freddo

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

func (f *Freddo) UpdateApp(c web.C, w http.ResponseWriter, req *http.Request) {
	appName := c.URLParams["appname"]

	app, ok := f.Apps[appName]
	if !ok {
		log.Printf("Could not find app: %s", appName)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Content Type"))
		return
	}

	body := new(bytes.Buffer)
	_, err := io.Copy(body, req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	signature := req.Header.Get("X-Hub-Signature")
	ok, err = app.HmacEqual(body.Bytes(), signature)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = app.QueueUpdate()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Update queue is full."))
		return
	}

	w.Write([]byte("OK"))
}
