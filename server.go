package main

import (
	"errors"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type configServer struct {
	data      map[string]*Config // izigrava bazu podataka
	dataGroup map[string]*Group
}

func (cs *configServer) createConfigHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := createId()
	rt.Id = id
	cs.data[id] = rt
	renderJSON(w, rt)
}

func (ts *configServer) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allTasks := []*Config{}
	for _, v := range ts.data {
		allTasks = append(allTasks, v)
	}

	renderJSON(w, allTasks)
}

func (ts *configServer) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := ts.data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (cs *configServer) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if v, ok := cs.data[id]; ok {
		delete(cs.data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

func (cs *configServer) createGroupHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	group, err := decodeGroup(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := createId()
	group.Id = id
	cs.dataGroup[id] = group
	renderJSON(w, group)
}

func (cs *configServer) AddConfigToGroup(w http.ResponseWriter, req *http.Request) {
	groupId := mux.Vars(req)["groupId"]
	id := mux.Vars(req)["id"]
	task, ok := cs.data[id]
	group, ook := cs.dataGroup[groupId]
	if !ok || !ook {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	group.Configs = append(group.Configs, *task)
	cs.dataGroup[groupId] = group

	return
}

func (cs *configServer) getAllGroupsHandler(w http.ResponseWriter, req *http.Request) {
	allGroups := []*Group{}
	for _, v := range cs.dataGroup {
		allGroups = append(allGroups, v)
	}

	renderJSON(w, allGroups)
}

func (cs *configServer) getGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := cs.dataGroup[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (cs *configServer) delGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	_, ok := cs.dataGroup[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	delete(cs.dataGroup, id)
}

func (cs *configServer) delConfigFromGroupHandler(w http.ResponseWriter, req *http.Request) {
	groupId := mux.Vars(req)["groupId"]
	id := mux.Vars(req)["id"]
	group, ok := cs.dataGroup[groupId]
	if !ok {
		err := errors.New("group not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	for i, config := range group.Configs {
		if config.Id == id {
			group.Configs = append(group.Configs[:i], group.Configs[i+1:]...)
			cs.dataGroup[groupId] = group
			return
		}
	}

	err := errors.New("config not found in group")
	http.Error(w, err.Error(), http.StatusNotFound)
	return
}
