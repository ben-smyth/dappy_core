package rest

import (
	"dappy_core/yaml"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func getApplications(w http.ResponseWriter, r *http.Request) {

	var config yaml.YAMLConfig
	config.SetFolderLocation("/home/bsmyth/GIT/dappy_core/index/")
	response, err := yaml.GetApplicationsFromYAML(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func getApplicaton(w http.ResponseWriter, r *http.Request) {
	application := mux.Vars(r)["name"]

	var config yaml.YAMLConfig
	config.SetFolderLocation("/home/bsmyth/GIT/dappy_core/index/")
	response, err := yaml.GetApplicationFromYAML(config, application)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
