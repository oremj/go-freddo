package freddo

import (
	"log"
	"net/http"

	"github.com/codegangsta/martini"
)

func (f *Freddo) UpdateApp(params martini.Params, res http.ResponseWriter) string {
	appName := params["app_name"]

	app, ok := f.Apps[appName]
	if !ok {
		log.Printf("Could not find app: %s", appName)
		res.WriteHeader(http.StatusNotFound)
		return ""
	}

	go app.Update()

	return "OK"
}
