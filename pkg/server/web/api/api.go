// Copyright 2019 Squeeze Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"net/http"
	log "github.com/golang/glog"
	"github.com/agile6v/squeeze/pkg/util"
	"github.com/agile6v/squeeze/pkg/server/web/controllers"
)

type AppAPI struct {
	MasterAddr string
	HTTPAddr   string
	LocalAddr  string
}

func (api *AppAPI) Init() {
	http.HandleFunc("/", api.Index)
	http.HandleFunc("/api/create", api.create)
	http.HandleFunc("/api/delete", api.delete)
	http.HandleFunc("/api/search", api.search)
	http.HandleFunc("/api/list", api.list)
	http.HandleFunc("/api/start", api.start)
	http.HandleFunc("/api/stop", api.stop)
	http.HandleFunc("/api/callback", api.callback)
}

func (api *AppAPI) Index(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, "")
}

func (api *AppAPI) create(w http.ResponseWriter, r *http.Request) {
	// Read body
	task := &controllers.CreateTask{}
	body, err := util.ReadBody(r, task)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = task.Handle(body)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, nil)
}

func (api *AppAPI) delete(w http.ResponseWriter, r *http.Request) {
	// Read body
	task := &controllers.GenericTask{}
	_, err := util.ReadBody(r, task)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = task.Delete()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, nil)
}

func (api *AppAPI) search(w http.ResponseWriter, r *http.Request) {

}

func (api *AppAPI) list(w http.ResponseWriter, r *http.Request) {
	tasks, err := controllers.ListTask()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, tasks)
}

func (api *AppAPI) start(w http.ResponseWriter, r *http.Request) {
	// Read body
	task := &controllers.GenericTask{}
	_, err := util.ReadBody(r, task)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = task.Start(api.MasterAddr, api.LocalAddr)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, nil)
}

func (api *AppAPI) stop(w http.ResponseWriter, r *http.Request) {
	// Read body
	task := &controllers.GenericTask{}
	_, err := util.ReadBody(r, task)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = task.Stop(api.MasterAddr)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithError(w, http.StatusOK, "")
}

func (api *AppAPI) callback(w http.ResponseWriter, r *http.Request) {
	body, err := util.ReadBody(r, nil)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Infof("Recv results: %s", body)

	err = controllers.HandleCallback(body)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, nil)
}