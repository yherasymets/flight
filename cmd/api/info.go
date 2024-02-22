package main

import (
	"net/http"
)

func (a *app) infoHandler(w http.ResponseWriter, r *http.Request) {
	js := wrapJson{
		"status": "available",
		"information": map[string]string{
			"environment": a.config.env,
			"version":     version,
		},
	}

	err := a.sendJSON(w, http.StatusOK, js, nil)
	if err != nil {
		a.serverErrorResp(w, r, err)
	}
}
