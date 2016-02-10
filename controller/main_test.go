package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
//	"time"

	"github.com/shipyard/shipyard/controller/api"
	"github.com/shipyard/shipyard/utils"
	"github.com/shipyard/shipyard/auth/builtin"
	"github.com/shipyard/shipyard/controller/manager"
//	"io"
//	"strings"
)

var (
	serverShipyard *httptest.Server
	serverSwarm    *httptest.Server
//	shipyardApi    *api.Api
)

func init() {
	 mockApi()
//	if err := a.Run(); err != nil {
//		panic(err)
//	}
//
//	shipyardApi = a

	serverShipyard = httptest.NewServer(api.ApiRouter)
	serverSwarm = httptest.NewServer(api.SwarmRouter)
}

func mockApi() {
	rethinkdbAddr := "rethinkdb:28015"
	rethinkdbDatabase := "shipyard"
	rethinkdbAuthKey := ""
	disableUsageInfo := false

	dockerUrl := "tcp://swarm:2375"
	tlsCaCert := ""
	tlsCert := ""
	tlsKey := ""
	allowInsecure := false

	client, err := utils.GetClient(dockerUrl, tlsCaCert, tlsCert, tlsKey, allowInsecure)
	if err != nil {
		panic(err)
	}

	// default to builtin auth
	authenticator := builtin.NewAuthenticator("defaultshipyard")

	controllerManager, err := manager.NewManager(rethinkdbAddr, rethinkdbDatabase, rethinkdbAuthKey, client, disableUsageInfo, authenticator)
	if err != nil {
		panic(err)
	}

	apiConfig := api.ApiConfig{
		ListenAddr:         ":8080",
		Manager:            controllerManager,
		AuthWhiteListCIDRs: nil,
		EnableCORS:         false,
		AllowInsecure:      false,
		TLSCACertPath:      "",
		TLSCertPath:        "",
		TLSKeyPath:         "",
	}

	shipyardApi, err := api.NewApi(apiConfig)
	if err != nil {
		panic(err)
	}


		if err := shipyardApi.Run(); err != nil {
			panic(err)
		}
}

func TestIntegration(t *testing.T) {
	res, err := http.Get("http://localhost:8080" + "/api/accounts")
	if err != nil {
		fmt.Printf("heres the error: %s", err)
		panic(err)
	}

	fmt.Printf("Here is the response to ping: %s", res)
}