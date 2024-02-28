package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/yherasymets/flight/models"
)

func (a *app) craeteFlight(w http.ResponseWriter, r *http.Request) {
	flight := new(models.Flight)
	if err := a.readJSON(w, r, &flight); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	if err := a.repo.Create(flight); err != nil {
		a.serverErrorResp(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/flight/%s", flight.ID.String()))
	err := a.sendJSON(w, http.StatusCreated, wrapJson{"flight": flight}, headers)
	if err != nil {
		a.serverErrorResp(w, r, err)
	}
}

func (a *app) getFlihgt(w http.ResponseWriter, r *http.Request) {
	id := a.getIDfromQuery(r)
	flight, err := a.repo.Get(id)
	if err != nil {
		a.serverErrorResp(w, r, err)
	}
	err = a.sendJSON(w, http.StatusOK, wrapJson{"data": flight}, nil)
	if err != nil {
		a.serverErrorResp(w, r, err)
	}
}

func (a *app) updateFlight(w http.ResponseWriter, r *http.Request) {
	id := a.getIDfromQuery(r)
	fligt := new(models.Flight)
	if err := a.readJSON(w, r, &fligt); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	prevFlight, err := a.repo.Get(id)
	if err != nil {
		a.serverErrorResp(w, r, err)
		return
	}
	fligt.CreatedAt = prevFlight.CreatedAt
	fligt.UpdatedAt = time.Now()
	fligt.ID = prevFlight.ID
	if err := a.repo.Update(id, fligt); err != nil {
		a.serverErrorResp(w, r, err)
	}
	err = a.sendJSON(w, http.StatusOK, wrapJson{"data": fligt}, nil)
	if err != nil {
		a.serverErrorResp(w, r, err)
	}
}

func (a *app) deleteFlihgt(w http.ResponseWriter, r *http.Request) {
	id := a.getIDfromQuery(r)
	err := a.repo.Delete(id)
	if err != nil {
		a.serverErrorResp(w, r, err)
	}
}

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
