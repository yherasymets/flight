package main

import (
	"fmt"
	"net/http"
)

func (a *app) logError(r *http.Request, err error) {
	a.logger.Println(err)
}

func (a *app) errResp(w http.ResponseWriter, r *http.Request, statusCode int, message interface{}) {
	wrap := wrapJson{"error": message}
	err := a.sendJSON(w, statusCode, wrap, nil)
	if err != nil {
		a.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *app) serverErrorResp(w http.ResponseWriter, r *http.Request, err error) {
	a.logError(r, err)
	message := "the server encountered an error and was unable to process your request"
	a.errResp(w, r, http.StatusInternalServerError, message)
}

func (a *app) notFoundResp(w http.ResponseWriter, r *http.Request) {
	message := "resource could not be found"
	a.errResp(w, r, http.StatusNotFound, message)
}

func (a *app) methodNotAllowedResp(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	a.errResp(w, r, http.StatusMethodNotAllowed, message)
}

func (a *app) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.errResp(w, r, http.StatusBadRequest, err.Error())
}

func (a *app) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	a.errResp(w, r, http.StatusUnprocessableEntity, errors)
}
