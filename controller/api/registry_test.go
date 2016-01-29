package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shipyard/shipyard/auth"
	"github.com/shipyard/shipyard/controller/mock_test"
	"github.com/stretchr/testify/assert"
)

// ENDPOINTS TO TEST

//func (a *Api) registries(w http.ResponseWriter, r *http.Request)
//
//func (a *Api) addRegistry(w http.ResponseWriter, r *http.Request)
//
//func (a *Api) registry(w http.ResponseWriter, r *http.Request)
//
//func (a *Api) removeRegistry(w http.ResponseWriter, r *http.Request)
//
//func (a *Api) repositories(w http.ResponseWriter, r *http.Request)
//
//func (a *Api) repository(w http.ResponseWriter, r *http.Request)
//
//func (a *Api) deleteRepository(w http.ResponseWriter, r *http.Request)
//
//func (a *Api) inspectRepository(w http.ResponseWriter, r *http.Request)

// TEST CASES (Take a look at accounts_test for an example on how test handlers)

func TestApiGetRegistries(t *testing.T) {

}

func TestApiAddRegistry(t *testing.T) {

}

func TestApiGetRegistry(t *testing.T) {

}

func TestApiRemoveRegistry(t *testing.T) {

}

func TestApiGetRepositories(t *testing.T) {

}

func TestApiGetRepository(t *testing.T) {

}

func TestApiDeleteRepository(t *testing.T) {

}

func TestApiInspectRepository(t *testing.T) {

}