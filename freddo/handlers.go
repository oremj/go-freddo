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
	app.Lock()
	defer app.Unlock()

	out, err := app.RunScript()
	if err != nil {
		log.Printf("Failed: %s", app.Script)
		log.Print("output: ", string(out))
		log.Print(err)
		res.WriteHeader(http.StatusBadRequest)
		return "Script failed."
	}

	return "OK"
}
