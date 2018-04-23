package api

import (
	"net/http"
	"os"

	"github.com/aziule/simple-logs-gui/listener"
)

// HandleOpenLocalFile starts tailing a local log file
func (api *Api) HandleOpenLocalFile(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Query().Get("path")

	if path == "" {
		api.writeError(w, "Missing parameter \"path\"", 400)
		return
	}

	if _, err := os.Stat(path); err != nil {
		api.writeError(w, "Could not open file or file does not exist", 400)
		return
	}

	err := listener.ListenToFile(path, listener.LocalListeningStrategy)

	if err != nil {
		api.writeError(w, err.Error(), 500)
		return
	}

	api.writeJson(w, map[string]interface{}{
		"status": "ok",
	})
}

// HandleOpenFile starts tailing a remote log file
func (api *Api) HandleOpenRemoteFile(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Query().Get("path")

	if path == "" {
		api.writeError(w, "Missing parameter \"path\"", 400)
		return
	}

	err := listener.ListenToFile(path, listener.RemoteListeningStrategy)

	if err != nil {
		api.writeError(w, err.Error(), 500)
		return
	}

	api.writeJson(w, map[string]interface{}{
		"status": "ok",
	})
}
