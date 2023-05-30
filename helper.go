package main

import (
	"encoding/json"
	"io"
	"main/configstore"
	"net/http"

	"github.com/google/uuid"
)

func decodeBody(r io.Reader) (*configstore.Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var cfg configstore.Config
	if err := dec.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func decodeGroup(r io.Reader) (*configstore.CfGroup, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var g configstore.CfGroup
	if err := dec.Decode(&g); err != nil {
		return nil, err
	}
	return &g, nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId() string {
	return uuid.New().String()
}
