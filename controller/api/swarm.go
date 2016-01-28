package api

import (
	"bytes"
	"net/http"
	"net/url"
	"encoding/json"
	"encoding/base64"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type AuthConfig struct {
	Username		string		`json:"username"`
	Password		string		`json:"password"`
	Auth			string		`json:"auth"`
	Email			string		`json:"email"`
}

func authConfigB64(username, password, auth, email string) string {
	log.Debugf("username: %s, pass: %s", username, password)

	authConfig := AuthConfig{
		Username: "admin",
		Password: "admin",
		Auth: "",
		Email: "",
	}

	authJson, err := json.Marshal(authConfig)
	if err != nil {
		return ""
	}

	authB64 := base64.StdEncoding.EncodeToString(authJson)

	return authB64
}

func registryHost(imageName string) string  {
	// Image format [REGISTRY/][USERNAME/]NAME[:TAG]
	if strings.Count(imageName, "/") == 2 {
		registry := strings.SplitN(imageName, "/", 2)[0]
		return registry
	} else {
		return ""
	}
}

func (a *Api) swarmRedirect(w http.ResponseWriter, req *http.Request) {
	var err error
	req.URL, err = url.ParseRequestURI(a.dUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Debugf("In swarm redirect")
	if imageName := req.Header.Get("Reg-Image-Name"); imageName != "" {
		log.Debugf("Reg-Image-Name detected!")
		if registryHost := registryHost(imageName); registryHost != "" {
			log.Debugf("image name contains registry! here it is %s", "https://"+registryHost)
			if registry, err := a.manager.RegistryByAddress("https://" + registryHost); err == nil {
				log.Debugf("setting auth header!")
				req.Header.Set("X-Registry-Auth", authConfigB64(registry.Username, registry.Password, "", ""))
			} else if registry, err := a.manager.RegistryByAddress("http://" + registryHost); err == nil {
				log.Debugf("setting auth header!")
				req.Header.Set("X-Registry-Auth", authConfigB64(registry.Username, registry.Password, "", ""))
			}
		}
	}

	a.fwd.ServeHTTP(w, req)
}

type proxyWriter struct {
	Body       *bytes.Buffer
	Headers    *map[string][]string
	StatusCode *int
}

func (p proxyWriter) Header() http.Header {
	return *p.Headers
}
func (p proxyWriter) Write(data []byte) (int, error) {
	return p.Body.Write(data)
}
func (p proxyWriter) WriteHeader(code int) {
	*p.StatusCode = code
}
