package api

import (
	"encoding/json"
	"net/http"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/shipyard/shipyard"
	"fmt"
)

func (a *Api) dockerLogin(w http.ResponseWriter, r *http.Request) {
	var registry *shipyard.Registry
	if err := json.NewDecoder(r.Body).Decode(&registry); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: see if there's a client library that can do docker logins
	loginCmd := fmt.Sprintf("sudo docker login -u %s -p %s -e none %s", registry.Username, registry.Password, registry.Addr)
	log.Infof("Exec Docker Login='%s'", loginCmd)

	if out, err := exec.Command(loginCmd).Output(); err != nil {
		log.Errorf("error executing docker login: %s", err)
		log.Errorf("%s: %s", loginCmd, out)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Infof("Successfully logged into registry=%s", registry.Addr)
	w.WriteHeader(http.StatusNoContent)
}